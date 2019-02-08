package delivery_test


import (	
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/siddhiparekh11/GoChallenge/delivery"
	"github.com/siddhiparekh11/GoChallenge/models"
	"github.com/gorilla/mux"
	"context"
	"strings"
	"strconv"
	"bytes"
)

/*
	Test file for Player Handler. Using httptest package to test the routes and handlers. 
*/


//fake struct to implement IPlayer interface
type fakePlayerContr struct {}


//fake GetPlayers implementation
func (f fakePlayerContr)GetPlayers(ctx context.Context) ([] *models.Player, *models.QError) {
	players := make([] *models.Player,0)
	players = append(players,&models.Player{1001,"John",400,4})
	players = append(players,&models.Player{1002,"Tom",210,4})
	return players,nil
}

//fake GetPlayer implementation
func (f fakePlayerContr)GetPlayer(ctx context.Context,playerId int) (*models.Player,*models.QError){
	player := &models.Player {1001,"John",400,4}
	return player,nil
}

//fake CreatePlayer implementation
func (f fakePlayerContr)CreatePlayer(ctx context.Context, player models.Player) (bool,*models.QError){
	return true,nil
}


//CreatePlayer handler test
func TestCreatePlayer(t *testing.T){
	player := models.Player{1001,"John",400,4}
	jsonPlayer, _ := json.Marshal(player)
	req, err := http.NewRequest("POST","/api/createplayer",bytes.NewBuffer(jsonPlayer))
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	rou := delivery.NewPlayerHandler(mux.NewRouter(),nil,fakePlayerContr{})
	rou.Mux.ServeHTTP(rec,req)
	assert.Equal(t,200,rec.Code,"Ok response is expected")
	res := strings.Replace(rec.Body.String(), "\n", "", -1)
	assert.Equal(t,"true",res,"Response is incorrect.")
}


//GetPlayer handler test
func TestGetPlayer(t *testing.T) {
  	id:= 1001
  	req, err := http.NewRequest("GET","/api/player/" + strconv.Itoa(id),nil)
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	rou := delivery.NewPlayerHandler(mux.NewRouter(),nil,fakePlayerContr{})
	rou.Mux.ServeHTTP(rec,req)
	assert.Equal(t,200,rec.Code,"Ok response is expected")
	str := strings.Replace(rec.Body.String(), "\n", "", -1)
	tstStr := `{"IdPlayer":1001,"Name":"John","ChipsAmount":400,"TotalNoChips":4}`
	assert.Equal(t,tstStr,str,"Response is incorrect.")
}


//GetPlayers handler test
func TestGetPlayers(t *testing.T) {
	req, err := http.NewRequest("GET","/api/players",nil)
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	rou := delivery.NewPlayerHandler(mux.NewRouter(),nil,fakePlayerContr{})
	rou.Mux.ServeHTTP(rec,req)
	assert.Equal(t,200,rec.Code,"Ok response is expected")
	str := strings.Replace(rec.Body.String(), "\n", "", -1)
	tstStr := `[{"IdPlayer":1001,"Name":"John","ChipsAmount":400,"TotalNoChips":4},{"IdPlayer":1002,"Name":"Tom","ChipsAmount":210,"TotalNoChips":4}]`
	assert.Equal(t,tstStr,str,"Response is incorrect.")
}