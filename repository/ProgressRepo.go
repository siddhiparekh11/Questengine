package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"context"
	"database/sql"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
	"log"
)

// Progress Repository. Contains Sql calls.
//1).UpdateProgress - update milestone and quest details for a particular player
//2).GetPlayerState - will fetch the last quest and milestone the player has completed
//3).GetAllQuestProgress - will fetch all the records belonging to the player and the current game


type ProgressRepository struct {
    Conn *sql.DB
}


//function return the IProgress interface
func NewProgressRepository(conn *sql.DB) (interfaces.IProgress) {
		return &ProgressRepository {conn}
}


//function to Update the quest progress of the player 
func (progressRepo *ProgressRepository) UpdateProgress(ctx context.Context, progress models.Progress) (*models.Progress,*models.QError) {
	log.Println("I am called from game repo update progress.")
	insert, err := progressRepo.Conn.Prepare("insert into Progress(idGame,idQuest,idPlayer,questPointsEarned,totalQuestComPercent,lastMilestoneInd,createdTimestamp) values(?,?,?,?,?,?,?)")
	if err != nil {
		return nil,&models.QError{Where: "ProgressRepo/UpdateProgress", What: err.Error()}
	}
	_, err = insert.Exec(progress.GameId,progress.QuestId,progress.PlayerId,progress.QuestPointsEarned,progress.TotQuestCompPer,progress.LastMilestoneIndex,progress.CreatedTimestamp)
	if err != nil {
		return nil,&models.QError{Where: "ProgressRepo/UpdateProgress", What: err.Error()}
	}
	p:= &progress
	return p,nil
}

//function to get the last completed quest and milestone
func (progressRepo *ProgressRepository) GetPlayerState(ctx context.Context,playerId int) (*models.Progress, *models.QError) {
	log.Println("I am called from game repo GetPlayerState.")
	query := "SELECT idGame,idQuest,idPlayer,questPointsEarned,totalQuestComPercent,lastMilestoneInd,createdTimestamp FROM Progress where idPlayer=? order by createdTimestamp desc limit 1"
	rows,err := progressRepo.Conn.QueryContext(ctx,query,playerId)
	if err!=nil {
		return nil,&models.QError{Where: "ProgressRepo/GetPlayerState", What: err.Error()}
	}
	defer rows.Close()
	progress := make([] *models.Progress,0)
	for rows.Next() {		
		p := new(models.Progress)
		err = rows.Scan(&p.GameId,&p.QuestId,&p.PlayerId,&p.QuestPointsEarned,&p.TotQuestCompPer,&p.LastMilestoneIndex,&p.CreatedTimestamp)		
		if err!=nil {
		return nil,&models.QError{Where: "ProgressRepo/GetPlayerState", What: err.Error()}
		}
		progress = append(progress,p)	
	}
	if len(progress) == 0 {
		return nil, &models.QError{Where: "ProgressRepo/GetPlayerState", What: "No such Progress Record"}
	}
	log.Println(progress[0])
	return progress[0],nil
}

//function to get the quest beloning to the player and current game
func (progressRepo *ProgressRepository)GetAllQuestProgress(ctx context.Context, playerId int, gameId int) ([]*models.Progress,*models.QError){
	log.Println("I am called from GetAllQuestProgress.")
	query := "SELECT idGame,idQuest,idPlayer,questPointsEarned,totalQuestComPercent,lastMilestoneInd,createdTimestamp FROM Progress where idPlayer=? and idGame=? order by createdTimestamp"
	rows,err := progressRepo.Conn.QueryContext(ctx,query,playerId,gameId)
	if err!=nil {
		return nil,&models.QError{Where: "ProgressRepo/GetPlayerState", What: err.Error()}
	}
	defer rows.Close()
	progress := make([] *models.Progress,0)
	for rows.Next() {		
		p := new(models.Progress)
		err = rows.Scan(&p.GameId,&p.QuestId,&p.PlayerId,&p.QuestPointsEarned,&p.TotQuestCompPer,&p.LastMilestoneIndex,&p.CreatedTimestamp)		
		if err!=nil {
		return nil,&models.QError{Where: "ProgressRepo/GetPlayerState", What: err.Error()}
		}
		progress = append(progress,p)	
	}
	if len(progress) == 0 {
		return nil, &models.QError{Where: "ProgressRepo/GetAllQuestProgress", What: "No such Progress Record"}
	}
	log.Println(len(progress))
	return progress,nil
}


//Sql implementation of this function is not required. It is implemented in Progress controller
func (progressRepo *ProgressRepository) ConstructPlayerStateStruct(progress *models.Progress) (*models.PlayerState) {
	return nil
}
