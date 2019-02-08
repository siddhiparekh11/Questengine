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


//GetPlayers controller test
func TestGetPlayers(t *testing.T){

	mockPlayerRepo := new(mocks.IPlayer)
	mockPlayers := []*models.Player{
			&models.Player {
				PlayerId: 1001,
				NamePlayer: "John",
				ChipsAmount : 400,
				TotalNoOfChips: 4,
			},
			&models.Player {
				PlayerId: 1002,
				NamePlayer: "Tom",
				ChipsAmount : 200,
				TotalNoOfChips: 2,
			},
	}	
	pc := controller.NewPlayerController(mockPlayerRepo)
	t.Run("success", func(t *testing.T) {
		mockPlayerRepo.On("GetPlayers", mock.Anything).Return(mockPlayers, nil).Once()
		playerChan := make(chan controller.PlayerChannel)
		go func(){
			players, err := pc.GetPlayers(context.TODO())
			chanObj := new(controller.PlayerChannel)
			chanObj.Players = players
			chanObj.Player = nil
			chanObj.Error = err
			chanObj.IsCreated = false
			playerChan<-(*chanObj)
		}()		
		players := make([] *models.Player,0)
		var err *models.QError
		for c := range playerChan {	
			players = c.Players
			err = c.Error
			break
		}	
		close(playerChan)
		require.Nil(t,err)
		assert.Len(t,players,len(mockPlayers))
		mockPlayerRepo.AssertExpectations(t)		
	})

	t.Run("error", func(t *testing.T) {
		mockPlayerRepo.On("GetPlayers", mock.Anything).Return(nil, &models.QError{}).Once()
		playerChan := make(chan controller.PlayerChannel)
		go func(){
			players, err := pc.GetPlayers(context.TODO())
			chanObj := new(controller.PlayerChannel)
			chanObj.Players = players
			chanObj.Player = nil
			chanObj.Error = err
			chanObj.IsCreated = false
			playerChan<-(*chanObj)
		}()		
		players := make([] *models.Player,0)
		var err *models.QError
		for c := range playerChan {	
			players = c.Players
			err = c.Error
			break
		}	
		close(playerChan)
		assert.NotEmpty(t,err)
		assert.Len(t,players,0)
		mockPlayerRepo.AssertExpectations(t)		
	})
}

//GetPlayer controller test
func TestGetPlayer(t *testing.T){
	mockPlayerRepo := new(mocks.IPlayer)
	mockPlayer := &models.Player {
				PlayerId: 1001,
				NamePlayer: "John",
				ChipsAmount : 400,
				TotalNoOfChips: 4,
			}	
	pc := controller.NewPlayerController(mockPlayerRepo)
	t.Run("success", func(t *testing.T) {
		mockPlayerRepo.On("GetPlayer", mock.Anything,mock.AnythingOfType("int")).Return(mockPlayer, nil).Once()
		playerChan := make(chan controller.PlayerChannel)
		go func(){
			player, err := pc.GetPlayer(context.TODO(),1001)
			chanObj := new(controller.PlayerChannel)
			chanObj.Players = nil
			chanObj.Player = player
			chanObj.Error = err
			chanObj.IsCreated = false
			playerChan<-(*chanObj)
		}()		
		var player *models.Player
		var err *models.QError
		for c := range playerChan {	
			player = c.Player
			err = c.Error
			break
		}	
		close(playerChan)
		require.Nil(t,err)
		assert.Equal(t,player,mockPlayer)
		mockPlayerRepo.AssertExpectations(t)		
	})

	t.Run("error", func(t *testing.T) {
		mockPlayerRepo.On("GetPlayer", mock.Anything,mock.AnythingOfType("int")).Return(nil, &models.QError{}).Once()
		playerChan := make(chan controller.PlayerChannel)
		go func(){
			player, err := pc.GetPlayer(context.TODO(),1001)
			chanObj := new(controller.PlayerChannel)
			chanObj.Players = nil
			chanObj.Player = player
			chanObj.Error = err
			chanObj.IsCreated = false
			playerChan<-(*chanObj)
		}()		
		var err *models.QError
		for c := range playerChan {	
			err = c.Error
			break
		}	
		close(playerChan)
		assert.NotEmpty(t,err)
		mockPlayerRepo.AssertExpectations(t)		
	})
}


//CreatePlayer controller test
func TestCreatePlayer(t *testing.T) {
	mockPlayerRepo := new(mocks.IPlayer)
	player := models.Player {
				PlayerId: 1001,
				NamePlayer: "John",
				ChipsAmount : 400,
				TotalNoOfChips: 4,
			}
	gc := controller.NewPlayerController(mockPlayerRepo)
	t.Run("success", func(t *testing.T) {
		mockPlayerRepo.On("CreatePlayer", mock.Anything,mock.AnythingOfType("models.Player")).Return(true, nil).Once()
		createChan := make(chan controller.PlayerChannel)
		go func(){
			flag, err := gc.CreatePlayer(context.TODO(),player)
			chanObj := new(controller.PlayerChannel)
			chanObj.Players = nil
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
		mockPlayerRepo.AssertExpectations(t)		
	})

	t.Run("error", func(t *testing.T) {
		mockPlayerRepo.On("CreatePlayer", mock.Anything,mock.AnythingOfType("models.Player")).Return(false, &models.QError{}).Once()
		createChan := make(chan controller.PlayerChannel)
		go func(){
			flag, err := gc.CreatePlayer(context.TODO(),player)
			chanObj := new(controller.PlayerChannel)
			chanObj.Players = nil
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
		mockPlayerRepo.AssertExpectations(t)		
	})
}
