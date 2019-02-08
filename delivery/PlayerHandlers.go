package delivery


import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/siddhiparekh11/GoChallenge/models"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"strconv"	
)

//Handlers for player apis


//Custom interface for player handlers
type IPlayerHandlers interface {
	GetPlayers(w http.ResponseWriter, r *http.Request) 
	GetPlayer(w http.ResponseWriter, r *http.Request)
	CreatePlayer(w http.ResponseWriter, r *http.Request)
}


type PlayerHandler struct {
	Mux *mux.Router
	conn *sql.DB
	IPlayerHandlers
	pContr interfaces.IPlayer
}


//function initialize all the routes to the player apis and return PlayerHandler obj
func NewPlayerHandler(m *mux.Router,con *sql.DB, pCon interfaces.IPlayer) (*PlayerHandler) {
	playerHandler := &PlayerHandler {
		Mux : m,
		conn: con,
		pContr: pCon,
	}
	playerHandler.Mux.HandleFunc("/api/players",playerHandler.GetPlayers).Methods("GET")
	playerHandler.Mux.HandleFunc("/api/player/{playerid}",playerHandler.GetPlayer).Methods("GET")
	playerHandler.Mux.HandleFunc("/api/createplayer",playerHandler.CreatePlayer).Methods("POST")
	return playerHandler
}


//GetPlayers handler
func (playerHandler *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	players, err := playerHandler.pContr.GetPlayers(context.Background())
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}else{
		json.NewEncoder(w).Encode(players)
	}
}

//GetPlayer handler
func (playerHandler *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	log.Println(vars)
	playerid,e := strconv.Atoi(vars["playerid"])
	if e!=nil {
		json.NewEncoder(w).Encode(e)
	}else{
		author, err := playerHandler.pContr.GetPlayer(context.Background(),playerid)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}else{
			json.NewEncoder(w).Encode(author)
		}
	}
}

//CreatePlayer handler
func (playerHandler *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	decoder := json.NewDecoder(r.Body)
	var player models.Player
	e := decoder.Decode(&player)
	if e != nil {
		json.NewEncoder(w).Encode(e)
	}else{
		flag, err := playerHandler.pContr.CreatePlayer(context.Background(),player)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}else{
			json.NewEncoder(w).Encode(flag)
		}
	}	
}

