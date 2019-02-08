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

	Game controller test. Used mocking framework.

*/

var config models.Config 


//GetGame controller test. Implemented subtests.
func TestGetGames(t *testing.T){
	mockGameRepo := new(mocks.IGame)
	mockGames := make([] *models.Game,0)
	mockGames = append(mockGames,&models.Game{101,"BlackJack","11,12"})
	mockGames = append(mockGames,&models.Game{102,"Pokemon","12,13"})
	gc := controller.NewGameController(mockGameRepo,config)
	t.Run("success", func(t *testing.T) {
		mockGameRepo.On("GetGames", mock.Anything).Return(mockGames, nil).Once()
		gamesChan := make(chan controller.GameChannel)
		go func(){
			games, err := gc.GetGames(context.TODO())
			chanObj := new(controller.GameChannel)
			chanObj.Games = games
			chanObj.Error = err
			chanObj.IsCreated = false
			gamesChan<-(*chanObj)
		}()		
		games := make([] *models.Game,0)
		var err *models.QError
		for c := range gamesChan {		
				games = c.Games
				err = c.Error
				break
		}
		close(gamesChan)
		require.Nil(t,err)
		assert.Len(t,games,len(mockGames))
		mockGameRepo.AssertExpectations(t)		
	})

	t.Run("error", func(t *testing.T) {
		mockGameRepo.On("GetGames", mock.Anything).Return(nil, &models.QError{}).Once()
		gamesChan := make(chan controller.GameChannel)
		go func(){
			games, err := gc.GetGames(context.TODO())
			chanObj := new(controller.GameChannel)
			chanObj.Games = games
			chanObj.Error = err
			chanObj.IsCreated = false
			gamesChan<-(*chanObj)
		}()		
		games := make([] *models.Game,0)
		var err *models.QError
		for c := range gamesChan {		
				games = c.Games
				err = c.Error
				break
		}
		close(gamesChan)
		assert.NotEmpty(t,err)
		assert.Len(t,games,0)
		mockGameRepo.AssertExpectations(t)		
	})
}


//CreateGame controller test. Implemented subtests.
func TestCreateGame(t *testing.T) {
	mockGameRepo := new(mocks.IGame)
	game := models.Game{101,"BlackJack","11,12"}
	gc := controller.NewGameController(mockGameRepo,config)
	t.Run("success", func(t *testing.T) {
		mockGameRepo.On("CreateGame", mock.Anything,mock.AnythingOfType("models.Game")).Return(true, nil).Once()
		createChan := make(chan controller.GameChannel)
		go func(){
			flag, err := gc.CreateGame(context.TODO(),game)
			chanObj := new(controller.GameChannel)
			chanObj.Games = nil
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
		mockGameRepo.AssertExpectations(t)		
	})

	t.Run("error", func(t *testing.T) {
		mockGameRepo.On("CreateGame", mock.Anything,mock.AnythingOfType("models.Game")).Return(false, &models.QError{}).Once()
		createChan := make(chan controller.GameChannel)
		go func(){
			flag, err := gc.CreateGame(context.TODO(),game)
			chanObj := new(controller.GameChannel)
			chanObj.Games = nil
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
		mockGameRepo.AssertExpectations(t)		
	})
}


//SetGameId controller test. Implemented subtests.
func TestSetGameId(t *testing.T){

	mockGameRepo := new(mocks.IGame)
	mockGame := &models.Game {
				GameId: 101,
				GameName: "BlackJack",
				QuestOrder:"1,2,3",
			}	
	
	gc := controller.NewGameController(mockGameRepo,config)
	t.Run("success", func(t *testing.T) {
		mockGameRepo.On("SetGameId", mock.Anything,mock.AnythingOfType("int")).Return(mockGame, nil).Once()
		gameChan := make(chan controller.GameChannel)
		go func(){
			game, err := gc.SetGameId(context.TODO(),101)
			chanObj := new(controller.GameChannel)
			chanObj.Games = nil
			chanObj.Game = game
			chanObj.Error = err
			chanObj.IsCreated = false
			gameChan<-(*chanObj)
		}()		
		var game *models.Game
		var err *models.QError
		for c := range gameChan {	
			game = c.Game
			err = c.Error
			break
		}	
		close(gameChan)
		require.Nil(t,err)
		assert.Equal(t,game,mockGame)
		mockGameRepo.AssertExpectations(t)		
	})

	t.Run("error", func(t *testing.T) {
		mockGameRepo.On("SetGameId", mock.Anything,mock.AnythingOfType("int")).Return(nil, &models.QError{}).Once()
		gameChan := make(chan controller.GameChannel)
		go func(){
			game, err := gc.SetGameId(context.TODO(),101)
			chanObj := new(controller.GameChannel)
			chanObj.Games = nil
			chanObj.Game = game
			chanObj.Error = err
			chanObj.IsCreated = false
			gameChan<-(*chanObj)
		}()		
		var err *models.QError
		for c := range gameChan {	
			err = c.Error
			break
		}	
		close(gameChan)
		assert.NotEmpty(t,err)
		mockGameRepo.AssertExpectations(t)		
	})
}
