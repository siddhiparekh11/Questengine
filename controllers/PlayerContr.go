package controller


import (	
	"context"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
	"log"
)

type PlayerController struct {
	PRepo interfaces.IPlayer
}


//Player controller - uses go routines

//go routine uses PlayerChannel communicate
type PlayerChannel struct {
	Players []*models.Player
	Error  *models.QError
	Player *models.Player
	IsCreated bool
}


//function return object of IPlayer interface
func NewPlayerController(pRepo interfaces.IPlayer) interfaces.IPlayer {
	return &PlayerController { PRepo: pRepo }
}


//PlayerController implements GetPlayers method. Call GetPlayers in Player repository
func (playerContr *PlayerController) GetPlayers(ctx context.Context) ([] *models.Player, *models.QError) {
	playersChan := make(chan PlayerChannel)	
	go func() {
		players,err := playerContr.PRepo.GetPlayers(ctx)
		chanObj := new(PlayerChannel)
		chanObj.Player = nil
		chanObj.IsCreated = false
		if err!=nil {
			chanObj.Players = nil
			chanObj.Error = err
		}else{
			chanObj.Players = players
			chanObj.Error = nil
		}
		playersChan <- *chanObj
	}()
	players := make([] *models.Player,0)
	var err *models.QError
	for c := range playersChan {	
			log.Println("I am called")	
			players = c.Players
			err = c.Error
			break
	}	
	close(playersChan)	
	if err!=nil {
		return nil, &models.QError{Where: "PlayerContr/GetPlayers", What: err.Error()}
	}
	return players, nil
}

//Playercontroller implements GetPlayer method. Call GetPlayer in player repository
func (playerContr *PlayerController) GetPlayer(ctx context.Context,playerId int) (*models.Player,*models.QError) {
	playersChan := make(chan PlayerChannel)	
	go func() {
		player,err := playerContr.PRepo.GetPlayer(ctx,playerId)
		chanObj := new(PlayerChannel)
		chanObj.Players = nil
		chanObj.IsCreated = false
		if err!=nil {
			chanObj.Player = nil
			chanObj.Error = err
		}else{
			chanObj.Player = player
			chanObj.Error = nil
		}
		playersChan <- *chanObj
	}()
	var player *models.Player
	var err *models.QError
	for c := range playersChan {	
			log.Println("I am called")	
			player = c.Player
			err = c.Error
			break
	}
	close(playersChan)
	if err!=nil {
		return nil, &models.QError{Where: "PlayerContr/GetPlayers", What: err.Error()}
	}
	return player, nil
}

//PlayerController implements CreatePlayer method. Call CreatePlayer in player repository
func (playerContr *PlayerController) CreatePlayer(ctx context.Context,player models.Player) (bool, *models.QError) {
	createChan := make(chan PlayerChannel)
	go func() {		
		isCreated, err := playerContr.PRepo.CreatePlayer(ctx,player)
		chanObj := new(PlayerChannel)
		chanObj.Players = nil
		chanObj.Player = nil
		if err!=nil {
			chanObj.Error = err
			chanObj.IsCreated = false		
		}else{
			chanObj.IsCreated = isCreated
			chanObj.Error = nil	
		}		
		createChan<-*chanObj
	}()
	var err *models.QError
	var isCreated bool
	for c := range createChan {	
			log.Println("I am called")	
			isCreated = c.IsCreated
			err = c.Error
			break
	}
	close(createChan)
	if err!=nil {
		return false, &models.QError{Where: "PlayerContr/CreatePlayer", What: err.Error()}
	}
	return isCreated, nil
}