package controller


import (	
	"context"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
	"log"
	"strconv"
	"strings"
	"time"
)

//Process controller - uses go routines for extensive processing

type ProgressController struct {
	PrRepo interfaces.IProgress
	QRepo interfaces.IQuest
	PRepo interfaces.IPlayer
	GRepo interfaces.IGame	
}


//ProgressChannel to communicate between go routines
type ProgressChannel struct {
	Progress *models.Progress
	ProgressArr []*models.Progress
	Error *models.QError	
}

var config models.Config

//Progress Controller returns an Object of IProgress interface
func NewProgressController(prRepo interfaces.IProgress, qRepo interfaces.IQuest, pRepo interfaces.IPlayer, gRepo interfaces.IGame,confg models.Config) interfaces.IProgress {	
	config = confg
	return &ProgressController { PrRepo: prRepo, QRepo: qRepo, PRepo: pRepo, GRepo: gRepo}
}

//Progress controller implements GetAllQuestProgress method. It gets all the progress pertaining to a player playing a game configured in config file.
func (progressContr *ProgressController)GetAllQuestProgress(ctx context.Context, playerId int, gameId int) ([]*models.Progress,*models.QError){
	progressChan := make(chan ProgressChannel)		
	go func() {
		progarr,err := progressContr.PrRepo.GetAllQuestProgress(ctx,playerId,gameId)
		chanObj := new(ProgressChannel)
		if err!=nil {
				chanObj.ProgressArr = nil
				chanObj.Error = err
		}else{
				chanObj.ProgressArr = progarr
				chanObj.Error = nil
		}
		progressChan <- *chanObj
	}()
	var progarr []*models.Progress
	var err *models.QError
	for c := range progressChan {	
			log.Println("I am called")	
			progarr = c.ProgressArr
			err = c.Error
			break
	}	
	close(progressChan)	
	if err!=nil {
		return nil, &models.QError{Where: "ProgressContr/GetAllQuestProgress", What: err.Error()}
	}
	log.Println("from controller")
	log.Println(len(progarr))
	return progarr, nil
}


//Progress controller implements GetPlayerState method. It gets the lastest quest and milestone completed by the player.
func (progressContr *ProgressController) GetPlayerState(ctx context.Context,playerId int) (*models.Progress, *models.QError) {
	progressChan := make(chan ProgressChannel)	
	go func() {
		progress,err := progressContr.PrRepo.GetPlayerState(ctx,playerId)
		chanObj := new(ProgressChannel)
		if err!=nil {
				chanObj.Progress = nil
				chanObj.Error = err
		}else{
				chanObj.Progress = progress
				chanObj.Error = nil
		}
		progressChan <- *chanObj

	}()
	var progress *models.Progress
	var err *models.QError
	for c := range progressChan {	
			log.Println("I am called")	
			progress = c.Progress
			err = c.Error
			break
	}	
	close(progressChan)	
	if err!=nil {
		return nil, &models.QError{Where: "ProgressContr/GetPlayerState", What: err.Error()}
	}
	log.Println("from controller")
	log.Println(progress)
	return progress, nil
}

//Progress controller implements Update Progress method. It registers the next milestone completed by the player in the current quest
func (progressContr *ProgressController) UpdateProgress(ctx context.Context,progress models.Progress) (*models.Progress, *models.QError) {
	isValidPlayer := make(chan bool,1)
	isValidLevel := make(chan bool,1)
	qChan := make(chan models.Quest,1)
	log.Println("ChipAmountBet")
	log.Println(progress.ChipAmountBet)
	go verifyPlayer(ctx,progressContr.PRepo,isValidPlayer,progress.PlayerId,progress.ChipAmountBet)	
	go verifyPlayerLevel(ctx,progressContr.QRepo,isValidLevel,qChan,progress.QuestId)	
	plyFlag := <-isValidPlayer
	lvlFlag := <-isValidLevel
	quest := <-qChan
	close(isValidPlayer)
	close(isValidLevel)
	close(qChan)
	var qFlag bool
   	if lvlFlag && plyFlag {
		valQuest := make(chan bool, 1)
		go doesQuestBelongToGame(ctx,progressContr.GRepo,valQuest,progress.QuestId)
		qFlag = <-valQuest
   	}else{
   		return nil,&models.QError{Where: "ProgressContr/UpdateProgress", What: "Player level or PlayerId is incorrect."}
   	}
   	if !qFlag {
   	return nil,&models.QError{Where: "ProgressContr/UpdateProgress", What: "Quest doesn't belong to the game."}
   	}	
	if plyFlag && lvlFlag && qFlag {
		errChan := make(chan string)		
		prg := make(chan *models.Progress)
		go prepareProgressObject(ctx,progressContr.PrRepo,quest.MilestonesOrder,errChan,prg)
		prg <-&progress
		e := <-errChan
		<- prg
		/*log.Println("progress")
		log.Println(tprg)
		var t models.Progress
		t = (*tprg)
		progress = t*/	
		log.Println(progress)	
		if e!="" {
			return nil, &models.QError{Where: "ProgressContr/UpdateProgress", What: e} 
		}		
		progressChan := make(chan ProgressChannel)		
		go func() {						
			p, err := progressContr.PrRepo.UpdateProgress(ctx,progress)
			chanObj := new(ProgressChannel)
			if err!=nil {
				chanObj.Progress = nil
				chanObj.Error = err
			}else{
				chanObj.Progress = p
				chanObj.Error = nil
			}
			progressChan <- *chanObj
		}()	
		var err *models.QError
		for p := range progressChan {
			err = p.Error
			break
		}
		close(progressChan)
		if err!=nil {
			return nil,&models.QError{Where: "ProgressContr/UpdateProgress", What: err.Error()}
		}
	}else{
		return nil,&models.QError{Where: "ProgressContr/UpdateProgress", What: "Either the player level or Player id is incorrect OR Player doesn't have sufficient amount to bet."}
	}
	log.Println("I am getting called")
	return &progress,nil
}

//function verifies if the player whose record needs to be updated is a valid player. Go routine in Update progress calls this function.
func verifyPlayer(ctx context.Context,PRepo interfaces.IPlayer,isValidPlayer chan bool,PlayerId int,ChipAmountBet int) {
	player,err := PRepo.GetPlayer(ctx,PlayerId)
	log.Println("go routine get player")
	log.Println(player)
	chanObj := new(PlayerChannel)
	chanObj.Player = player
	chanObj.Error = err
	log.Println(err)
	if chanObj.Error == nil {
		log.Println(chanObj.Player.ChipsAmount)
		log.Println(ChipAmountBet)
		if chanObj.Player.ChipsAmount >= ChipAmountBet {
				isValidPlayer <- true
		}else{
				isValidPlayer <- false 
		}			
	}else{
			log.Println("I am called from go routine 1")
			isValidPlayer <- false
	}
}

//function verifies if the playerlevel/quest is a valid quest. Go routine in the Update progress calls this function
func verifyPlayerLevel(ctx context.Context,QRepo interfaces.IQuest,isValidLevel chan bool,qChan chan models.Quest,QuestId int){
	quest,err := QRepo.GetQuest(ctx,QuestId)
	log.Println("go routine get quest")
	log.Println(quest)
	log.Println(err)
	if err == nil {
		isValidLevel <- true
		qChan <- (*quest)
	}else{
		log.Println("I am called from go routine 2")
		isValidLevel <- false
		qChan<-models.Quest{0,"",""}
	}	
}

//function verifies if the quest belongs to the current game configured in the config file. Go routine in the Update progress calls this function.
func doesQuestBelongToGame(ctx context.Context,GRepo interfaces.IGame,valQuest chan bool,questId int){
	game, err := GRepo.SetGameId(ctx,config.GameId)
	log.Println("go routine get quest")
	log.Println(game)
	log.Println(err)
	if err != nil {
				valQuest<-false
	}else{
		flag := isValidQuest(game.QuestOrder,questId)
		if flag {
			valQuest<-true
		}else{
			valQuest<-false
		}
	}
}


//function calculates questpointsearned, milestones array, totalquestpercent completed. Go routine in Update progress calls this function.
func prepareProgressObject(ctx context.Context,PrRepo interfaces.IProgress,questMilestoneOrder string ,errChan chan string,prog chan *models.Progress){
	progress := <- prog
	log.Println("Player Id")
	log.Println(progress.PlayerId)
	p, err := PrRepo.GetPlayerState(ctx,progress.PlayerId)
	if err!=nil {
				errChan<-err.Error()
				prog<-progress
	}else{
			m, err := nextMilestone(p.LastMilestoneIndex)				
			if err!=nil {
						errChan<-err.Error()
						prog <- progress
			}else{
					s := splitString(questMilestoneOrder)
					sInd,_ := strconv.Atoi(s[0])
					eInd,_ := (strconv.Atoi(s[len(s)-1]))
					progarr,err := PrRepo.GetAllQuestProgress(ctx,progress.PlayerId,config.GameId)
					if err!=nil {
						errChan<-""
					}
					if (eInd>=m.Ind && sInd<=m.Ind) {
						progress.LastMilestoneIndex = m.Ind
						if p.QuestId==progress.QuestId {
							progress.TotQuestCompPer = calPercentage(len(s),p.TotQuestCompPer)
						}else{
							progress.TotQuestCompPer = calPercentage(len(s),0)
						}
						progress.LastMilestone = (*m)						
						calQuestPointsAccumulates(progress)
						if p.QuestId==progress.QuestId {
							progress.QuestPointsEarned += p.QuestPointsEarned
						}
						setMilestonesArray(progress,0,m.Ind,progarr)
						progress.GameId = config.GameId
						progress.CreatedTimestamp = time.Now()
						errChan<-""
						prog <- progress
					}else{
						errChan<-"You have completed all Milestones in this quest"
						prog <- progress
					}

			}
	}
}




//function prepares the data for valid json output. Not all fields of progress obj are exported by the api.
func (progressContr *ProgressController) ConstructPlayerStateStruct(progress *models.Progress) (*models.PlayerState) {
	var plyState models.PlayerState
	plyState.LastMilestoneIndex = progress.LastMilestoneIndex
	plyState.TotalQuestPercentCompleted = progress.TotQuestCompPer
	plyState.PlayerLevel= progress.QuestId
	return &plyState
}

//function is a part of another function 'DoesQuestBelongstoGame'. The quest order in the game is comma separated list.
func isValidQuest(questOrder string,questId int) bool {
	s := splitString(questOrder)
	for i:=0;i<len(s);i++{
		num,err := strconv.Atoi(s[i])
		if err!=nil {
			return false
		}
		if num==questId {
			return true;
		}
	}
	return false
}

//function prepares an array of current milestones and milestones completed in the past
func setMilestonesArray(progress *models.Progress,sInd int,eInd int, progarr []*models.Progress){
	mArr := make([]models.CusMilestone,0)
	log.Println(sInd)
	log.Println(eInd)
	log.Println(len(progarr))
	j:=0
	for i:=sInd ;i<eInd;i++ {
		var mil models.CusMilestone
		var chipsAwr int
		if(i==eInd-1){
			mil.Ind = config.Milestones[i].Ind
			if (config.Milestones[i].PointsNeededToWin>progress.QuestPointsEarned) {
				mil.ChipsAwarded=0
			}else{
				mil.ChipsAwarded=config.Milestones[i].ChipsAwarded
			}			
			progress.LastMilestone.ChipsAwarded = mil.ChipsAwarded
		}else{
			mil.Ind = progarr[j].LastMilestoneIndex
			if (config.Milestones[i].PointsNeededToWin>progarr[j].QuestPointsEarned) {
				chipsAwr=0
			}else{
				chipsAwr=config.Milestones[i].ChipsAwarded
			}
			mil.ChipsAwarded = chipsAwr // chips awarded state is not stored in database so all the previous milestone reward will be zero in the output
			j++
		}
		mArr = append(mArr,mil)
	}
	progress.Milestones = mArr
}

//function calculates the percentage of quest completed - quest completed is different than game completed percent
func calPercentage(totMilInQuest int,currPer int) int{
	newPer := currPer + (100/totMilInQuest)
	return newPer
}

//function splits an incoming string and returns a string slice
func splitString(str string) ([]string) {
	s := strings.Split(str,",")
	return s
}

//function calculates the next milestones index
func nextMilestone(lastMilestoneInd int) (*models.CusMilestone,*models.QError) {
	ind := lastMilestoneInd + 1
	if (ind>=1 && ind<=5){
		log.Println("index")
		log.Println(ind)
		mil := new(models.CusMilestone)
		mil.Ind = config.Milestones[ind-1].Ind
		return mil,nil
	}
	return nil, &models.QError{Where: "ProgressContr/nextMilestone", What: "You have reached the limit of configured milestones."}
}

//function calculates total quest points accumulated - it calculates quest points pertaining to a particular milestone 
func calQuestPointsAccumulates(progress *models.Progress){
	rate := config.Ratefrombet
	bonus := config.Levelbonusrate
	level:= progress.QuestId
	total := (float64(progress.ChipAmountBet) * rate) + (float64(level) * bonus)
	progress.QuestPointsEarned = int(total)
	log.Println(progress.QuestPointsEarned)
}