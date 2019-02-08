package controller


import (	
	"context"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
)

type QuestController struct {
	QRepo interfaces.IQuest
}


//QuestChannel is used to communicate between go routines
type QuestChannel struct {
	Quests []*models.Quest
	Error *models.QError
	Quest *models.Quest
	IsCreated bool
}


//function returns an object of IQuest interface
func NewQuestController(qRepo interfaces.IQuest) interfaces.IQuest {
	return &QuestController { QRepo: qRepo }
}


//Quest controller implements GetQuests method.
func (questContr *QuestController) GetQuests(ctx context.Context) ([] *models.Quest, *models.QError) {
	questsChan := make(chan QuestChannel)	
	go func() {
		quests,err := questContr.QRepo.GetQuests(ctx)
		chanObj := new(QuestChannel)
		chanObj.Quest = nil
		chanObj.IsCreated = false
		if err!=nil {
			chanObj.Quests = nil
			chanObj.Error = err
		}else{
			chanObj.Quests = quests
			chanObj.Error = nil
		}
		questsChan <- *chanObj
	}()
	quests := make([] *models.Quest,0)
	var err *models.QError
	for c := range questsChan {	
			quests = c.Quests
			err = c.Error
			break
	}
	close(questsChan)
	if err!=nil {
		return nil, &models.QError{Where: "QuestContr/GetQuests", What: err.Error()}
	}	
	return quests, nil
}

//Quest controller implements GetQuest method
func (questContr *QuestController) GetQuest(ctx context.Context,questId int) (*models.Quest,*models.QError) {
	questsChan := make(chan QuestChannel)	
	go func() {
		quest,err := questContr.QRepo.GetQuest(ctx,questId)
		chanObj := new(QuestChannel)
		chanObj.Quests = nil
		chanObj.IsCreated = false
		if err!=nil {
				chanObj.Quest = nil
				chanObj.Error = err
		}else{
				chanObj.Quest = quest
				chanObj.Error = nil
		}
		questsChan <- *chanObj
	}()	
	var quest *models.Quest
	var err *models.QError	
	for c := range questsChan {	
			quest = c.Quest
			err = c.Error
			break
	}
	close(questsChan)
	if err!=nil {
		return nil, &models.QError{Where: "QuestContr/CreateQuest", What: err.Error()}
	}	
	return quest, nil
}

//Quest controller implements CreateQuest method
func (questContr *QuestController) CreateQuest(ctx context.Context,quest models.Quest) (bool,*models.QError) {
	createChan := make(chan QuestChannel)
	go func() {
		isCreated, err := questContr.QRepo.CreateQuest(ctx,quest)
		chanObj := new(QuestChannel)		
		chanObj.Quests = nil
		chanObj.Quest = nil
		if err !=nil {
				chanObj.IsCreated = false
				chanObj.Error = err
		}else{
				chanObj.IsCreated = isCreated
				chanObj.Error = nil
		}		
		createChan <- *chanObj
	}()
	var err *models.QError
	var isCreated bool
	for c := range createChan {	
			isCreated = c.IsCreated
			err = c.Error
			break
	}
	close(createChan)
	if err!=nil {
		return false, &models.QError{Where: "QuestContr/CreateQuest", What: err.Error()}
	}
	return isCreated, nil
}