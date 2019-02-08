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
	"time"
	"log"
)



var (
	mockQuestRepo *mocks.IQuest
	mockPlayerRepo *mocks.IPlayer
	mockProgressRepo *mocks.IProgress
	mockGameRepo *mocks.IGame
	mockProgress *models.Progress
	mockPlayer *models.Player
	mockQuest *models.Quest
	mockGame *models.Game
	mockProArr []*models.Progress
)

func init(){
	mockQuestRepo = new(mocks.IQuest)
	mockPlayerRepo = new(mocks.IPlayer)
	mockProgressRepo = new(mocks.IProgress)
	mockGameRepo = new (mocks.IGame)
	mockProgress = &models.Progress{			
				GameId: 101,
				QuestId: 11, 
				PlayerId: 1001,  
				QuestPointsEarned: 103,
				TotQuestCompPer: 100,
				LastMilestoneIndex: 1,
				LastMilestone: models.CusMilestone{1,100},
				Milestones: []models.CusMilestone{{1,100}},
				CreatedTimestamp: time.Now(),
				ChipAmountBet: 200,			
	}

	mockProArr = []*models.Progress {
		 &models.Progress{			
				GameId: 101,
				QuestId: 11, 
				PlayerId: 1001,  
				QuestPointsEarned: 103,
				TotQuestCompPer: 100,
				LastMilestoneIndex: 1,
				LastMilestone: models.CusMilestone{1,100},
				Milestones: []models.CusMilestone{{1,100}},
				CreatedTimestamp: time.Now(),
				ChipAmountBet: 200,			
	},	
	}
	mockPlayer = &models.Player {
				PlayerId: 1001,
				NamePlayer: "John",
				ChipsAmount : 400,
				TotalNoOfChips: 4,
	}
	mockQuest = &models.Quest {
				QuestId: 12,
				QuestName: "Level2",
				MilestonesOrder:"2,3",
	}
	mockGame = &models.Game {
				GameId: 101,
				GameName: "BlackJack",
				QuestOrder: "11,12,13",
	}
	config.Milestones = []models.Milestone{{1,150,100},{2,200,150},{3,250,250},{4,100,300},{5,300,350}}
	config.Ratefrombet = 0.5
	config.Levelbonusrate = 0.25
	config.GameId = 101
}

//GetPlayerState controller test
func TestGetPlayerState(t *testing.T){

	pr := controller.NewProgressController(mockProgressRepo,mockQuestRepo,mockPlayerRepo,mockGameRepo,config)
	t.Run("success", func(t *testing.T) {
		mockProgressRepo.On("GetPlayerState", mock.Anything,mock.AnythingOfType("int")).Return(mockProgress, nil).Once()
		progressChan := make(chan controller.ProgressChannel)
		go func(){
			progress, err := pr.GetPlayerState(context.TODO(),1001)
			chanObj := new(controller.ProgressChannel)
			chanObj.Progress = progress
			chanObj.Error = err
			progressChan<-(*chanObj)
		}()		
		var progress *models.Progress
		var err *models.QError

		for c := range progressChan {		
				progress = c.Progress
				err = c.Error
				break
		}
		close(progressChan)
		require.Nil(t,err)
		assert.Equal(t,progress,mockProgress)
		mockProgressRepo.AssertExpectations(t)		
	})

	t.Run("error", func(t *testing.T) {
		mockProgressRepo.On("GetPlayerState", mock.Anything,mock.AnythingOfType("int")).Return(nil, &models.QError{}).Once()
		progressChan := make(chan controller.ProgressChannel)
		go func(){
			progress, err := pr.GetPlayerState(context.TODO(),1001)
			chanObj := new(controller.ProgressChannel)
			chanObj.Progress = progress
			chanObj.Error = err
			progressChan<-(*chanObj)
		}()		
		var err *models.QError
		for c := range progressChan {		
				err = c.Error
				break
		}
		close(progressChan)
		assert.NotEmpty(t,err)
		mockProgressRepo.AssertExpectations(t)		
	})
}

//UpdateProgress controller test
func TestUpdateProgress(t *testing.T){
	pr:= controller.NewProgressController(mockProgressRepo,mockQuestRepo,mockPlayerRepo,mockGameRepo,config)
	progrss:= &models.Progress{			
				GameId: 101,
				QuestId: 12, 
				PlayerId: 1001,  
				QuestPointsEarned: 53,
				TotQuestCompPer: 50,
				LastMilestoneIndex: 2,
				LastMilestone: models.CusMilestone{2,0},
				Milestones: []models.CusMilestone{{1,0},{2,0}},
				CreatedTimestamp: time.Now(),
				ChipAmountBet: 100,			
	}
	

	t.Run("successplayer", func(t *testing.T){

		mockPlayerRepo.On("GetPlayer", mock.Anything,mock.AnythingOfType("int")).Return(mockPlayer, nil).Once()
		mockQuestRepo.On("GetQuest", mock.Anything,mock.AnythingOfType("int")).Return(mockQuest, nil).Once()
		mockGameRepo.On("SetGameId", mock.Anything,mock.AnythingOfType("int")).Return(mockGame,nil).Once()
		mockProgressRepo.On("GetPlayerState", mock.Anything,mock.AnythingOfType("int")).Return(mockProgress, nil).Once()
		mockProgressRepo.On("GetAllQuestProgress",mock.Anything,mock.AnythingOfType("int"),mock.AnythingOfType("int")).Return(mockProArr,nil).Once()
		mockProgressRepo.On("UpdateProgress",mock.Anything,mock.AnythingOfType("models.Progress")).Return(progrss,nil).Once()
		progressChan := make(chan controller.ProgressChannel)
		go func(){
			progress, err := pr.UpdateProgress(context.TODO(),(*progrss))
			chanObj := new(controller.ProgressChannel)
			chanObj.Progress = progress
			chanObj.Error = err
			progressChan<-(*chanObj)
		}()
		var err *models.QError
		var progress *models.Progress

		for p := range progressChan {
			progress = p.Progress
			err = p.Error
			break
		}
		close(progressChan)
		require.Nil(t,err)
		log.Println(progress)
		//assert.Equal(t,progress,progrss) //will not be equal because timestamp
		mockProgressRepo.AssertExpectations(t)	

	})


	//many test cases are possible - player, playerlevel, game invalid
	//quest doesnt belong to the current game
	//something went wrong with update progress
	t.Run("error", func(t *testing.T){

		mockPlayerRepo.On("GetPlayer", mock.Anything,mock.AnythingOfType("int")).Return(mockPlayer,nil).Once()
		mockQuestRepo.On("GetQuest", mock.Anything,mock.AnythingOfType("int")).Return(mockQuest, nil).Once()
		mockGameRepo.On("SetGameId", mock.Anything,mock.AnythingOfType("int")).Return(mockGame,nil).Once()
		mockProgressRepo.On("GetPlayerState", mock.Anything,mock.AnythingOfType("int")).Return(mockProgress, nil).Once()
		mockProgressRepo.On("GetAllQuestProgress",mock.Anything,mock.AnythingOfType("int"),mock.AnythingOfType("int")).Return(mockProArr,nil).Once()
		mockProgressRepo.On("UpdateProgress",mock.Anything,mock.AnythingOfType("models.Progress")).Return(nil,&models.QError{}).Once()
		progressChan := make(chan controller.ProgressChannel)
		go func(){
			progress, err := pr.UpdateProgress(context.TODO(),(*progrss))
			chanObj := new(controller.ProgressChannel)
			chanObj.Progress = progress
			chanObj.Error = err
			progressChan<-(*chanObj)
		}()
		var err *models.QError

		for p := range progressChan {
			err = p.Error
			break
		}
		close(progressChan)
		assert.NotEmpty(t,err)
		mockProgressRepo.AssertExpectations(t)	

	})
}

