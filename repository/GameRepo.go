
package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"context"
	"database/sql"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"	
)

// Game repository. Contains sql calls.
// 1). Create Game - will create a new game. Each game will have certain quest order.
// 2). GetGames - will return all the created games
// 3). SetGameId - in the present code the gameid comes from config file but this can be altered using this method


type GameRepository struct {
    Conn *sql.DB
}


//function return the IGame interface
func NewGameRepository(conn *sql.DB) (interfaces.IGame) {
		return &GameRepository { conn }
}


//function to create a game
func (gameRepo *GameRepository) CreateGame(ctx context.Context, game models.Game) (bool,*models.QError) {	
	insert, err := gameRepo.Conn.Prepare("insert into Games(nameGame,questOrder) values(?,?)")
	if err != nil {
		return false, &models.QError{Where: "GameRepo/CreateGame", What: err.Error()}
	}	
	_, err = insert.Exec(game.GameName,game.QuestOrder)
	if err != nil {
		return false, &models.QError{Where: "GameRepo/CreateGame", What: err.Error()}
	}
	return true,nil
}

//function to get all the games
func (gameRepo *GameRepository) GetGames(ctx context.Context) ([] *models.Game, *models.QError) {
	
	query := "SELECT idGames,nameGame,questOrder from Games"
	rows, err := gameRepo.Conn.QueryContext(ctx,query)
	if err!=nil {
		
		return nil,&models.QError{Where: "GameRepo/GetGames", What: err.Error()}
	}
	defer rows.Close()
	games := make([] *models.Game,0)
	for rows.Next() {		
		g := new(models.Game)
		err = rows.Scan(&g.GameId,&g.GameName,&g.QuestOrder)		
		if err!=nil {
		return nil,&models.QError{Where: "GameRepo/GetGames", What: err.Error()}
		}
		games = append(games,g)	
	}
	if len(games) == 0 {
		return nil, &models.QError{Where: "GameRepo/SetGameId", What: "No game records."}
	}
	return games,nil
}

//function to get details of a requested game - in the controller the new gameid is set in the config struct
func (gameRepo *GameRepository) SetGameId(ctx context.Context,gameId int)(*models.Game, *models.QError){
	query := "SELECT idGames,nameGame,questOrder from Games where idGames=?"
	rows, err := gameRepo.Conn.QueryContext(ctx,query,gameId)
	if err!=nil {		
		return nil,&models.QError{Where: "GameRepo/GetGames", What: err.Error()}
	}
	defer rows.Close()
	games := make([] *models.Game,0)
	for rows.Next() {		
		g := new(models.Game)
		err = rows.Scan(&g.GameId,&g.GameName,&g.QuestOrder)		
		if err!=nil {
		return nil,&models.QError{Where: "GameRepo/SetGameId", What: err.Error()}
		}
		games = append(games,g)	
	}
	if len(games) == 0 {
		return nil, &models.QError{Where: "GameRepo/SetGameId", What: "No such Game Id"}
	}
	return games[0],nil
}
