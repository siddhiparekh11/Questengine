package delivery


import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/siddhiparekh11/GoChallenge/models"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/controllers"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strconv"	
)

//Handlers for Game Apis


//custom game handlers interface
type IGameHandlers interface {
	GetGames(w http.ResponseWriter, r *http.Request)
	CreateGame(w http.ResponseWriter, r *http.Request)
}


type GameHandler struct {
	Mux *mux.Router
	conn *sql.DB
	IGameHandlers
	gContr interfaces.IGame //game controllers
}



//function initialize all the routes to the game apis and return gamehandler obj
func NewGameHandler(m *mux.Router,con *sql.DB, gCon interfaces.IGame) (*GameHandler) {
	gameHandler := &GameHandler {
		Mux : m,
		conn: con,
		gContr: gCon,
	}
	gameHandler.Mux.HandleFunc("/api/games",gameHandler.GetGames).Methods("GET")
	gameHandler.Mux.HandleFunc("/api/creategame",gameHandler.CreateGame).Methods("POST")
	gameHandler.Mux.HandleFunc("/api/setgameid/{gameid}",gameHandler.SetGameId).Methods("GET")
	return gameHandler
}


//GetGames handler
func (gameHandler *GameHandler) GetGames(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	games, err := gameHandler.gContr.GetGames(context.Background())
	if err != nil {
		json.NewEncoder(w).Encode(err)		
	}else{
		json.NewEncoder(w).Encode(games)
	}
}

//CreateGame handler
func (gameHandler *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")	
	decoder := json.NewDecoder(r.Body)
	var game models.Game
	e := decoder.Decode(&game)
	if e != nil {
		json.NewEncoder(w).Encode(e)
	}else{
		if !(controller.QuestOrderValidation(game.QuestOrder)) {
			json.NewEncoder(w).Encode("Invalid Input")
		}else{	
			flag, err := gameHandler.gContr.CreateGame(context.Background(),game)
			if err != nil {
				json.NewEncoder(w).Encode(err)
			}else{
				json.NewEncoder(w).Encode(flag)
			}
		}
	}	
}

//SetGameId handler
func (gameHandler *GameHandler) SetGameId(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	gameId, e := strconv.Atoi(vars["gameid"])
	if e != nil {
		json.NewEncoder(w).Encode(e)
	}else{
		game, err := gameHandler.gContr.SetGameId(context.Background(),gameId)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}else{
			json.NewEncoder(w).Encode(game)
		}
	}
}

