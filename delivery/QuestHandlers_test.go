package delivery_test


import (
	
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	//"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/siddhiparekh11/GoChallenge/delivery"
	"github.com/siddhiparekh11/GoChallenge/models"
	//"io/ioutil"
	"github.com/gorilla/mux"
	"context"
	"strings"
	"strconv"
	"bytes"

)

/*
	Test file for Quest Handler. Using httptest package to test the routes and handlers. 
*/


//fake struct to implement IQuest interface
type fakeQuestContr struct {}


//fake GetQuests implementation
func (f fakeQuestContr)GetQuests(ctx context.Context) ([] *models.Quest, *models.QError) {


	quests := make([] *models.Quest,0)
	quests = append(quests,&models.Quest{11,"level1","1,2"})
	quests = append(quests,&models.Quest{12,"level2","3,4"})

	return quests,nil

}

//fake GetQuest implementation
func (f fakeQuestContr)GetQuest(ctx context.Context,questId int) (*models.Quest,*models.QError){
	quest := &models.Quest {11,"level1","1,2"}
	return quest,nil
}

//fake CreateQuest implementation
func (f fakeQuestContr)CreateQuest(ctx context.Context, quest models.Quest) (bool,*models.QError){
	return true,nil
}

//CreateQuest handler test
func TestCreateQuest(t *testing.T){

	quest := models.Quest {11,"level1","1,2"}

	jsonQuest, _ := json.Marshal(quest)
	req, err := http.NewRequest("POST","/api/createquest",bytes.NewBuffer(jsonQuest))
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	rou := delivery.NewQuestHandler(mux.NewRouter(),nil,fakeQuestContr{})

	rou.Mux.ServeHTTP(rec,req)

	assert.Equal(t,200,rec.Code,"Ok response is expected")
	res := strings.Replace(rec.Body.String(), "\n", "", -1)

	assert.Equal(t,"true",res,"Response is incorrect.")

}


//GetQuest handler test
func TestGetQuest(t *testing.T) {

  	id:= 11
  	req, err := http.NewRequest("GET","/api/quest/" + strconv.Itoa(id),nil)
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	rou := delivery.NewQuestHandler(mux.NewRouter(),nil,fakeQuestContr{})

	rou.Mux.ServeHTTP(rec,req)

	assert.Equal(t,200,rec.Code,"Ok response is expected")

	str := strings.Replace(rec.Body.String(), "\n", "", -1)
	tstStr := `{"IdQuest":11,"Name":"level1","MilestonesOrder":"1,2"}`

	assert.Equal(t,tstStr,str,"Response is incorrect.")

}



//GetQuests handler test
func TestGetQuests(t *testing.T) {

	req, err := http.NewRequest("GET","/api/quests",nil)
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()

	rou := delivery.NewQuestHandler(mux.NewRouter(),nil,fakeQuestContr{})

	rou.Mux.ServeHTTP(rec,req)

	assert.Equal(t,200,rec.Code,"Ok response is expected")

	str := strings.Replace(rec.Body.String(), "\n", "", -1)
	tstStr := `[{"IdQuest":11,"Name":"level1","MilestonesOrder":"1,2"},{"IdQuest":12,"Name":"level2","MilestonesOrder":"3,4"}]`

	assert.Equal(t,tstStr,str,"Response is incorrect.")
	

}