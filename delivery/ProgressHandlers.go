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
	"strconv"
)

//Progress handlers apis


//custom interface for Progress handlers
type IProgressHandlers interface {
	UpdateProgress(w http.ResponseWriter, r *http.Request)
	GetPlayerState(w http.ResponseWriter, r *http.Request) 
}


type ProgressHandler struct {
	Mux *mux.Router
	conn *sql.DB
	IProgressHandlers
	prContr interfaces.IProgress
}


//function initialize all the routes and return ProgressHandler obj
func NewProgressHandler(m *mux.Router,con *sql.DB, prCon interfaces.IProgress) (*ProgressHandler) {
	progressHandler := &ProgressHandler {
		Mux : m,
		conn: con,
		prContr: prCon,
	}
	progressHandler.Mux.HandleFunc("/api/progress",progressHandler.UpdateProgress).Methods("POST")
	progressHandler.Mux.HandleFunc("/api/state/player/{playerid}",progressHandler.GetPlayerState).Methods("GET")
	return progressHandler
}


//GetPlayerState handler
func (progressHandler *ProgressHandler) GetPlayerState(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	playerid,e:= strconv.Atoi(vars["playerid"])	
	if e!=nil{
		json.NewEncoder(w).Encode(e)
	}else{
		progress, err := progressHandler.prContr.GetPlayerState(context.Background(),playerid)
		if err != nil {		
			json.NewEncoder(w).Encode(err)
		}else{
			plyState := progressHandler.prContr.ConstructPlayerStateStruct(progress)
			json.NewEncoder(w).Encode(plyState)
		}
	}
}

//UpdateProgress handler
func (progressHandler *ProgressHandler) UpdateProgress(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	decoder := json.NewDecoder(r.Body)
	var prog models.Prog
	e := decoder.Decode(&prog)
	if e != nil {
		json.NewEncoder(w).Encode(e)
	}else{
		var progress models.Progress	
		progress.PlayerId = prog.PlayerId
		progress.QuestId = prog.PlayerLevel
		progress.ChipAmountBet = prog.ChipAmountBet
		p, err := progressHandler.prContr.UpdateProgress(context.Background(),progress)		
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}else{
			var prg models.ProgressView
			prg.PlayerLevel = p.QuestId
			prg.QuestPointsEarned = p.QuestPointsEarned
			prg.Milestones= p.Milestones
			prg.TotQuestCompPer=p.TotQuestCompPer
			json.NewEncoder(w).Encode(prg)
		}
	}
}

