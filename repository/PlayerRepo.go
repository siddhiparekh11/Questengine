package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"context"
	"database/sql"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
)

//Player Repository. Contains Sql calls.
//1).Create Player will create a player. All Players will start with a certain chipsamount which they can use in any games.
//2).GetPlayers will get all the registered players. Currently they are not associated with a particular game
//3).GetPlayer will get detials of specific player.



type PlayerRepository struct {
    Conn *sql.DB
}


//function return the IPlayer interface
func NewPlayerRepository(conn *sql.DB) (interfaces.IPlayer) {
		return &PlayerRepository {conn}
}


//function to create player
func (playerRepo *PlayerRepository) CreatePlayer(ctx context.Context, player models.Player) (bool,*models.QError) {
	insert, err := playerRepo.Conn.Prepare("insert into Players(namePlayer,chipsAmount,totalNoChips) values(?,?,?)")
	if err != nil {
		return false,&models.QError{Where: "playerRepo/CreatePlayer", What: err.Error()}
	}	
	_, err = insert.Exec(player.NamePlayer,player.ChipsAmount,player.TotalNoOfChips)
	if err != nil {
		return false,&models.QError{Where: "playerRepo/CreatePlayer", What: err.Error()}
	}
	return true,nil
}

//function to get all the players
func (playerRepo *PlayerRepository) GetPlayers(ctx context.Context) ([] *models.Player, *models.QError) {
	query := "SELECT idPlayers,namePlayer,chipsAmount,totalNoChips from Players"
	rows,err := playerRepo.Conn.QueryContext(ctx,query)
	if err!=nil {		
		return nil,&models.QError{Where: "playerRepo/GetPlayers", What: err.Error()}
	}
	defer rows.Close()
	players := make([] *models.Player,0)
	for rows.Next() {		
		p := new(models.Player)
		err = rows.Scan(&p.PlayerId,&p.NamePlayer,&p.ChipsAmount,&p.TotalNoOfChips)		
		if err!=nil {
			return nil,&models.QError{Where: "playerRepo/GetPlayers", What: err.Error()}
		}
		players = append(players,p)
	
	}
	if len(players) == 0 {
		return nil, &models.QError{Where: "PlayerRepo/GetPlayer", What: "No player records"}
	}
	return players,nil

}

//function to get a particular player
func (playerRepo *PlayerRepository) GetPlayer(ctx context.Context, playerId int) (*models.Player, *models.QError) {
	query := "SELECT idPlayers,namePlayer,chipsAmount,totalNoChips from Players where idPlayers=?"
	rows,err := playerRepo.Conn.QueryContext(ctx,query,playerId)
	if err!=nil {
		return nil,&models.QError{Where: "playerRepo/GetPlayer", What: err.Error()}
	}
	defer rows.Close()
	players := make([] *models.Player,0)
	for rows.Next() {		
		p := new(models.Player)
		err = rows.Scan(&p.PlayerId,&p.NamePlayer,&p.ChipsAmount,&p.TotalNoOfChips)		
		if err!=nil {
		return nil,&models.QError{Where: "playerRepo/GetPlayer", What: err.Error()}
		}
		players = append(players,p)	
	}
	if len(players) == 0 {
		return nil, &models.QError{Where: "PlayerRepo/GetPlayer", What: "No such Player Id"}
	}
	return players[0],nil
}

