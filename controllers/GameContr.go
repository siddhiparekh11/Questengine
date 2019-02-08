package controller


import (	
	"context"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
	"log"
	"strings"
	"strconv"
)

// Game controller - uses go routines 


type GameController struct {
	GRepo interfaces.IGame
	config *models.Config
}


//go routines uses GameChannel to communicate
type GameChannel struct {
	Games []*models.Game
	Error *models.QError
	Game *models.Game
	IsCreated bool
}


//function return an object of IGame interface
func NewGameController(gRepo interfaces.IGame,confg models.Config) interfaces.IGame {
	return &GameController { GRepo: gRepo, config: &confg }
}


//GameController implementing GetGames method. Call GetGames in Game repository
func (gameContr *GameController) GetGames(ctx context.Context) ([] *models.Game, *models.QError) {	
	gamesChan := make(chan GameChannel)	
	go func() {
		games,err := gameContr.GRepo.GetGames(ctx)
		chanObj := new(GameChannel)
		if err!=nil {
			chanObj.Games = nil
			chanObj.Error = err
		}else {
			chanObj.Games = games
			chanObj.Error = nil
		}		
		chanObj.IsCreated = false
		gamesChan <- *chanObj
	}()
	games := make([] *models.Game,0)
	var err *models.QError
	for c := range gamesChan {	
			log.Println("I am called from range")	
			games = c.Games
			err = c.Error
			break
	}
	close(gamesChan)
	if err!=nil {
		return nil, &models.QError{Where: "GameContr/GetGames", What: err.Error()}
	}	
	return games, nil
}


//Gamecontroller implementing CreateGame method. Call CreateGame in Game repository
func (gameContr *GameController) CreateGame(ctx context.Context,game models.Game) (bool, *models.QError) {	
	createChan := make(chan GameChannel)
	go func() {
		isCreated, err := gameContr.GRepo.CreateGame(ctx,game)
		chanObj := new(GameChannel)		
		chanObj.Games = nil
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
			log.Println("I am called")	
			isCreated = c.IsCreated
			err = c.Error
			break
	}
	close(createChan)
	if err!=nil {
		return false, &models.QError{Where: "GameContr/CreateGames", What: err.Error()}
	}
	return isCreated, nil
}


//Gamecontroller implementing SetGameId method. Call SetGameId in Game repository
func (gameContr *GameController) SetGameId(ctx context.Context, gameId int) (*models.Game, *models.QError){
	gamesChan := make(chan GameChannel)
	go func() {
		game,err := gameContr.GRepo.SetGameId(ctx,gameId)
		chanObj := new(GameChannel)
		if err!=nil {
			chanObj.Game = nil
			chanObj.Error = err
		}else {
			chanObj.Game = game
			chanObj.Error = nil
		}		
		chanObj.IsCreated = false
		gamesChan <- *chanObj
	}()
	var game *models.Game
	var err *models.QError
	for c := range gamesChan {	
			log.Println("I am called from range")	
			game = c.Game
			err = c.Error
			break
	}
	close(gamesChan)
	if err!=nil {
		return nil, &models.QError{Where: "GameContr/SetGameId", What: err.Error()}
	}
	gameContr.config.GameId = game.GameId	
	return game, nil
}


//Game Handler and Quest handler calls this function to validate input format - avoids sql injection
func QuestOrderValidation(input string) bool {
	if input=="" {
		return false
	}
	s := strings.Split(input,",")
	for i:=0;i<len(s);i++{
		_,err := strconv.Atoi(s[i])
		if err!= nil {
			return false
		}
	}
	return true
}

