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

//Quest handler apis


//custom interface for Quest handlers
type IQuestHandlers interface {
	CreateQuest(w http.ResponseWriter, r *http.Request)
	GetQuests(w http.ResponseWriter, r *http.Request)
	GetQuest(w http.ResponseWriter, r * http.Request) 
}


type QuestHandler struct {
	Mux *mux.Router
	conn *sql.DB
	IQuestHandlers
	qContr interfaces.IQuest
}

//function initialize all the routes and return QuestHandler object
func NewQuestHandler(m *mux.Router,con *sql.DB, qCon interfaces.IQuest) (*QuestHandler) {
	questHandler := &QuestHandler {
		Mux : m,
		conn: con,
		qContr: qCon,
	}
	questHandler.Mux.HandleFunc("/api/quests",questHandler.GetQuests).Methods("GET")
	questHandler.Mux.HandleFunc("/api/quest/{questid}",questHandler.GetQuest).Methods("GET")
	questHandler.Mux.HandleFunc("/api/createquest",questHandler.CreateQuest).Methods("POST")
	return questHandler
}


//GetQuests handler
func (questHandler *QuestHandler) GetQuests(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	quests, err := questHandler.qContr.GetQuests(context.Background())
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}else{
		json.NewEncoder(w).Encode(quests)
	}
}

//GetQuest handler
func (questHandler *QuestHandler) GetQuest(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	questid,e := strconv.Atoi(vars["questid"])
	if e!=nil {
		json.NewEncoder(w).Encode(e)
	}else{
		quest, err := questHandler.qContr.GetQuest(context.Background(),questid)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}else{
			json.NewEncoder(w).Encode(quest)
		}
	}
}

//CreateQuest handler
func (questHandler *QuestHandler) CreateQuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	decoder := json.NewDecoder(r.Body)
	var quest models.Quest
	e := decoder.Decode(&quest)
	if e != nil {
		json.NewEncoder(w).Encode(e)
	}else{
		if !(controller.QuestOrderValidation(quest.MilestonesOrder)) {
			json.NewEncoder(w).Encode("Invalid Input")
		}else{
			flag, err := questHandler.qContr.CreateQuest(context.Background(),quest)
			if err != nil {
				json.NewEncoder(w).Encode(err)
			}else{
				json.NewEncoder(w).Encode(flag)
			}
		}
	}	
}

