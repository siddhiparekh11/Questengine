package controller_test


import (	
	"context"
	"testing"
	"github.com/siddhiparekh11/GoChallenge/models"
	"github.com/siddhiparekh11/GoChallenge/controllers"
	"github.com/siddhiparekh11/GoChallenge/interfaces/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

/*

	Quest controller test. Used mocking framework.

*/



//Get Quest controller test
func TestGetQuests(t *testing.T){

	mockQuestRepo := new(mocks.IQuest)
	mockQuests := []*models.Quest{
			&models.Quest {
				QuestId: 11,
				QuestName: "Level4",
				MilestonesOrder:"2,3",
			},
			&models.Quest {
				QuestId: 12,
				QuestName: "Level5",
				MilestonesOrder:"4,5",
			},
	}
	
	qc := controller.NewQuestController(mockQuestRepo)
	t.Run("success", func(t *testing.T) {
		mockQuestRepo.On("GetQuests", mock.Anything).Return(mockQuests, nil).Once()
		questChan := make(chan controller.QuestChannel)
		go func(){
			quests, err := qc.GetQuests(context.TODO())
			chanObj := new(controller.QuestChannel)
			chanObj.Quests = quests
			chanObj.Quest = nil
			chanObj.Error = err
			chanObj.IsCreated = false
			questChan<-(*chanObj)
		}()		
		quests := make([] *models.Quest,0)
		var err *models.QError
		for c := range questChan {	
			quests = c.Quests
			err = c.Error
			break
		}	
		close(questChan)
		require.Nil(t,err)
		assert.Len(t,quests,len(mockQuests))
		mockQuestRepo.AssertExpectations(t)
		
	})

	t.Run("error", func(t *testing.T) {
		mockQuestRepo.On("GetQuests", mock.Anything).Return(nil, &models.QError{}).Once()
		questChan := make(chan controller.QuestChannel)
		go func(){
			quests, err := qc.GetQuests(context.TODO())
			chanObj := new(controller.QuestChannel)
			chanObj.Quests = quests
			chanObj.Quest = nil
			chanObj.Error = err
			chanObj.IsCreated = false
			questChan<-(*chanObj)
		}()		
		quests := make([] *models.Quest,0)
		var err *models.QError
		for c := range questChan {	
			quests = c.Quests
			err = c.Error
			break
		}	
		close(questChan)
		assert.NotEmpty(t,err)
		assert.Len(t,quests,0)
		mockQuestRepo.AssertExpectations(t)		
	})
}

//GetQuest controller test
func TestGetQuest(t *testing.T){
	mockQuestRepo := new(mocks.IQuest)
	mockQuest := &models.Quest {
				QuestId: 15,
				QuestName: "Level5",
				MilestonesOrder:"4,5",
			}	
	qc := controller.NewQuestController(mockQuestRepo)
	t.Run("success", func(t *testing.T) {
		mockQuestRepo.On("GetQuest", mock.Anything,mock.AnythingOfType("int")).Return(mockQuest, nil).Once()
		questChan := make(chan controller.QuestChannel)
		go func(){
			quest, err := qc.GetQuest(context.TODO(),15)
			chanObj := new(controller.QuestChannel)
			chanObj.Quests = nil
			chanObj.Quest = quest
			chanObj.Error = err
			chanObj.IsCreated = false
			questChan<-(*chanObj)
		}()		
		var quest *models.Quest
		var err *models.QError
		for c := range questChan {	
			quest = c.Quest
			err = c.Error
			break
		}	
		close(questChan)
		require.Nil(t,err)
		assert.Equal(t,quest,mockQuest)
		mockQuestRepo.AssertExpectations(t)
		
	})

	t.Run("error", func(t *testing.T) {
		mockQuestRepo.On("GetQuest", mock.Anything,mock.AnythingOfType("int")).Return(nil, &models.QError{}).Once()
		questChan := make(chan controller.QuestChannel)
		go func(){
			quest, err := qc.GetQuest(context.TODO(),15)
			chanObj := new(controller.QuestChannel)
			chanObj.Quests = nil
			chanObj.Quest = quest
			chanObj.Error = err
			chanObj.IsCreated = false
			questChan<-(*chanObj)
		}()		
		var err *models.QError
		for c := range questChan {	
			err = c.Error
			break
		}	
		close(questChan)
		assert.NotEmpty(t,err)
		mockQuestRepo.AssertExpectations(t)
		
	})
}



//CreateQuest controller test
func TestCreateQuest(t *testing.T) {

	mockQuestRepo := new(mocks.IQuest)
	quest := models.Quest {
				QuestId: 15,
				QuestName: "Level5",
				MilestonesOrder:"4,5",
			}
	qc := controller.NewQuestController(mockQuestRepo)
	t.Run("success", func(t *testing.T) {
		mockQuestRepo.On("CreateQuest", mock.Anything,mock.AnythingOfType("models.Quest")).Return(true, nil).Once()
		createChan := make(chan controller.QuestChannel)
		go func(){
			flag, err := qc.CreateQuest(context.TODO(),quest)
			chanObj := new(controller.QuestChannel)
			chanObj.Quests = nil
			chanObj.Error = err
			chanObj.IsCreated = flag
			createChan<-(*chanObj)
		}()		
		var err *models.QError
		var isCreated bool
		for c := range createChan {		
			isCreated = c.IsCreated
			err = c.Error
			break
		}
		close(createChan)
		require.Nil(t,err)
		assert.Equal(t,true,isCreated)
		mockQuestRepo.AssertExpectations(t)
		
	})

	t.Run("error", func(t *testing.T) {
		mockQuestRepo.On("CreateQuest", mock.Anything,mock.AnythingOfType("models.Quest")).Return(false, &models.QError{}).Once()
		createChan := make(chan controller.QuestChannel)
		go func(){
			flag, err := qc.CreateQuest(context.TODO(),quest)
			chanObj := new(controller.QuestChannel)
			chanObj.Quests = nil
			chanObj.Error = err
			chanObj.IsCreated = flag
			createChan<-(*chanObj)
		}()		
		var err *models.QError
		var isCreated bool
		for c := range createChan {	
			isCreated = c.IsCreated
			err = c.Error
			break
		}
		close(createChan)
		assert.NotEmpty(t,err)
		assert.Equal(t,false,isCreated)
		mockQuestRepo.AssertExpectations(t)
		
	})

}
