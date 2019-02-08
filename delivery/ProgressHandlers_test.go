package delivery_test


import (
	
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/siddhiparekh11/GoChallenge/delivery"
	"github.com/siddhiparekh11/GoChallenge/models"
	//"io/ioutil"
	"github.com/gorilla/mux"
	"context"
	"strings"
	"strconv"
	"bytes"
	"time"

)

/*
	Test file for Progress Handler. Using httptest package to test the routes and handlers. 
*/

//fake struct to implement IProgress interface
type fakeProgressContr struct {}

//write test tables for different paths ...subtest need to be implemented

//fake GetPlayerState implemtation
func (f fakeProgressContr)GetPlayerState(ctx context.Context, playerId int) (*models.Progress,*models.QError){

	progress := &models.Progress {
		GameId: 101,
		QuestId: 11,  
		PlayerId: 1001,  
		QuestPointsEarned: 203, 
		TotQuestCompPer: 100,
		LastMilestoneIndex: 1,
		LastMilestone: models.CusMilestone{1,2}, 
		Milestones: []models.CusMilestone{{1,2},{2,1}}, 
		CreatedTimestamp: time.Now(), 
		ChipAmountBet : 30,
	}

	return progress,nil
	
}

//fake GetAllQuestProgress implemtation
func (f fakeProgressContr)GetAllQuestProgress(ctx context.Context, questId int, gameId int) ([]*models.Progress,*models.QError){
	return nil,nil
}

//fake UpdateProgress implementation
func (f fakeProgressContr)UpdateProgress(ctx context.Context, prgress models.Progress) (*models.Progress,*models.QError){
	
	fmt.Println("I am called from UpdateProgress")
	progress := &models.Progress {
		GameId: 101,
		QuestId: 11,  
		PlayerId: 1001,  
		QuestPointsEarned: 203, 
		TotQuestCompPer: 100,
		LastMilestoneIndex: 1,
		LastMilestone: models.CusMilestone{1,2}, 
		Milestones: []models.CusMilestone{{1,2},{2,1}}, 
		CreatedTimestamp: time.Now(), 
		ChipAmountBet : 30,
	}
	return progress,nil
}

//fake ConstructPlayerStateStruct implemtation
func (f fakeProgressContr)ConstructPlayerStateStruct(progress *models.Progress) (*models.PlayerState){
	
	plystate := &models.PlayerState{
		TotalQuestPercentCompleted: progress.TotQuestCompPer,
		LastMilestoneIndex: progress.LastMilestoneIndex,
	}

	return plystate
}

//UpdateProgress handler test
func TestUpdateProgress(t *testing.T){

	prog:= &models.Prog{1001,11,400}

	jsonProg, _ := json.Marshal(prog)
	req, err := http.NewRequest("POST","/api/progress",bytes.NewBuffer(jsonProg))
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	rou := delivery.NewProgressHandler(mux.NewRouter(),nil,fakeProgressContr{})

	rou.Mux.ServeHTTP(rec,req)

	assert.Equal(t,200,rec.Code,"Ok response is expected")
	fmt.Println(rec.Body)
	//jsonPro, _ := json.Marshal(rec.Body.String())
	res := strings.Replace(rec.Body.String(), "\n", "", -1)
	tstr:= `{"PlayerLevel":11,"QuestPointsEarned":203,"TotalQuestPercentCompleted":100,"MilestonesCompleted":[{"MilestoneIndex":1,"ChipsAwarded":2},{"MilestoneIndex":2,"ChipsAwarded":1}]}`

	assert.Equal(t,tstr,res,"Response is incorrect.")

}


//GetPlayerState handler test
func TestGetPlayerState(t *testing.T) {

	/*tt := []struct {
		name   string
		value  int
		ans    string
		err    string
	}{
		{name: "Valid input", value: 1001, ans: `{"PlayerLevel":0,"TotalQuestPercentCompleted":100,"LastMilestoneIndex":1}`},
		{name: "Invalid input", value: "hello", err: "Not a number"},
	}*/

  	id:= 1001
  	req, err := http.NewRequest("GET","/api/state/player/" + strconv.Itoa(id),nil)
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	rou := delivery.NewProgressHandler(mux.NewRouter(),nil,fakeProgressContr{})

	rou.Mux.ServeHTTP(rec,req)

	assert.Equal(t,200,rec.Code,"Ok response is expected")

	str := strings.Replace(rec.Body.String(), "\n", "", -1)
	tstStr := `{"PlayerLevel":0,"TotalQuestPercentCompleted":100,"LastMilestoneIndex":1}`

	assert.Equal(t,tstStr,str,"Response is incorrect.")

}



