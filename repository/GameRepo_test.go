package repository_test

import (
	repo "github.com/siddhiparekh11/GoChallenge/repository"
	"github.com/siddhiparekh11/GoChallenge/models"
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"context"
	"log"
)

/*
	Test file for Game Repo. Mocking the database and testing against predefined input. 
*/


//GetGames test
func TestGetGames (t *testing.T) {
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockGames := []models.Game{
			models.Game {
				GameId : 101,
				GameName: "Black Jack",
				QuestOrder : "11,12,13",
			},
			models.Game {
				GameId : 102,
				GameName: "Pokemon",
				QuestOrder: "13,14",
			},
	}
	rows := sqlmock.NewRows([]string{"idGames","nameGame","questOrder"}).
			AddRow(mockGames[0].GameId,mockGames[0].GameName,mockGames[0].QuestOrder).
			AddRow(mockGames[1].GameId,mockGames[1].GameName,mockGames[1].QuestOrder)
	query := "SELECT idGames,nameGame,questOrder from Games"	
	mock.ExpectQuery(query).WillReturnRows(rows)
	g := repo.NewGameRepository(db)
	games,err := g.GetGames(context.TODO())
	log.Println(err)
	require.Nil(t, err) // require is required to address custom error which implements Error interface
	assert.NotNil(t, games)
}


//SetGameId test
func TestSetGameId(t *testing.T){	
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gm := models.Game{
		GameId: 101,
		GameName:     "Las Vegas",
		QuestOrder:   "12,13",		
	}
	mockGames := []models.Game{
			models.Game {
				GameId : 101,
				GameName: "Las Vegas",
				QuestOrder : "12,13",
			},			
	}
	rows := sqlmock.NewRows([]string{"idGames","nameGame","questOrder"}).
			AddRow(mockGames[0].GameId,mockGames[0].GameName,mockGames[0].QuestOrder)
	query := "SELECT idGames,nameGame,questOrder from Games where idGames=\\?"	
	mock.ExpectQuery(query).WithArgs(gm.GameId).WillReturnRows(rows)
	g := repo.NewGameRepository(db)
	games,err := g.SetGameId(context.TODO(),gm.GameId)
	log.Println(err)
	require.Nil(t, err) // require is required to address custom error which implements Error interface
	assert.NotNil(t, games)	
}


//CreateGame test
func TestCreateGame(t *testing.T) {
	gm := models.Game{
		GameId: 101,
		GameName:     "Las Vegas",
		QuestOrder:   "12,13",		
	}	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}	
	defer db.Close()
	query := "insert into Games\\(nameGame,questOrder\\) values\\(\\?,\\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(gm.GameName,gm.QuestOrder).WillReturnResult(sqlmock.NewResult(105, 1))
	g := repo.NewGameRepository(db)
	flag,e := g.CreateGame(context.TODO(), gm)
	require.Nil(t, e)
	assert.Equal(t, flag, true)
}