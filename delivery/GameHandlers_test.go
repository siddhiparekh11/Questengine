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
	"bytes"
	"strconv"
)

/*
	Test file for Game Handler. Using httptest package to test the routes and handlers. 
*/


//fake Game Controller struct to implement IGame interface
type fakeGameContr struct {}

//fake GetGames implementation
func (f fakeGameContr)GetGames(ctx context.Context) ([] *models.Game, *models.QError) {
	games := make([] *models.Game,0)
	games = append(games,&models.Game{101,"BlackJack","11,12"})
	games = append(games,&models.Game{102,"Pokemon","12,13"})
	return games,nil
}

//fake SetGameId implementation
func (f fakeGameContr) SetGameId(ctx context.Context,gameId int) (*models.Game,*models.QError){
	game := &models.Game {
		GameId: 101,
		GameName: "BlackJack",
		QuestOrder: "11,12",
	}
	return game,nil
}

//fake CreateGame implementation
func (f fakeGameContr)CreateGame(ctx context.Context, game models.Game) (bool,*models.QError){
	return true,nil
}

//CreateGames handler test
func TestCreateGames(t *testing.T){
	game := models.Game {
		GameId: 101,
		GameName: "BlackJack",
		QuestOrder: "11,12",
	}
	jsonGame, _ := json.Marshal(game)
	req, err := http.NewRequest("POST","/api/creategame",bytes.NewBuffer(jsonGame))
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	rou := delivery.NewGameHandler(mux.NewRouter(),nil,fakeGameContr{})
	rou.Mux.ServeHTTP(rec,req)
	assert.Equal(t,200,rec.Code,"Ok response is expected")
	res := strings.Replace(rec.Body.String(), "\n", "", -1)
	assert.Equal(t,"true",res,"Response is incorrect.")
}

//SetGameId handler test
func TestSetGameId(t *testing.T){
	id:= 101
	req, err := http.NewRequest("GET","/api/setgameid/" + strconv.Itoa(id),nil)
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	rou := delivery.NewGameHandler(mux.NewRouter(),nil,fakeGameContr{})
	rou.Mux.ServeHTTP(rec,req)
	assert.Equal(t,200,rec.Code,"Ok response is expected")
	str := strings.Replace(rec.Body.String(), "\n", "", -1)
	tstStr := `{"IdGame":101,"Name":"BlackJack","QuestOrder":"11,12"}`
	assert.Equal(t,tstStr,str,"Response is incorrect.")
}



//GetGames handler test
func TestGetGames(t *testing.T) {
	req, err := http.NewRequest("GET","/api/games",nil)	
	if err != nil {
				t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	rou := delivery.NewGameHandler(mux.NewRouter(),nil,fakeGameContr{})
	rou.Mux.ServeHTTP(rec,req)
	assert.Equal(t,200,rec.Code,"Ok response is expected")
	str := strings.Replace(rec.Body.String(), "\n", "", -1)
	tstStr := `[{"IdGame":101,"Name":"BlackJack","QuestOrder":"11,12"},{"IdGame":102,"Name":"Pokemon","QuestOrder":"12,13"}]`
	assert.Equal(t,tstStr,str,"Response is incorrect.")
}