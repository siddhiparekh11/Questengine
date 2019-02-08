package repository_test

import (
	repo "github.com/siddhiparekh11/GoChallenge/repository"
	"github.com/siddhiparekh11/GoChallenge/models"
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"context"
)

/*
	Test file for Player Repo. Mocking the database and testing against predefined input. 
*/



//GetPlayers test
func TestGetPlayers (t *testing.T) {
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockPlayers := []models.Player{
			models.Player {
				PlayerId : 1001,
				NamePlayer: "John",
				ChipsAmount : 400,
				TotalNoOfChips: 4,
			},
			models.Player {
				PlayerId : 1002,
				NamePlayer: "Tom",
				ChipsAmount : 200,
				TotalNoOfChips: 2,
			},
	}
	rows := sqlmock.NewRows([]string{"idPlayers","namePlayer","chipsAmount","totalNoChips"}).
			AddRow(mockPlayers[0].PlayerId,mockPlayers[0].NamePlayer,mockPlayers[0].ChipsAmount,mockPlayers[0].TotalNoOfChips).
			AddRow(mockPlayers[1].PlayerId,mockPlayers[1].NamePlayer,mockPlayers[1].ChipsAmount,mockPlayers[1].TotalNoOfChips)
	query := "SELECT idPlayers,namePlayer,chipsAmount,totalNoChips from Players"	
	mock.ExpectQuery(query).WillReturnRows(rows)
	g := repo.NewPlayerRepository(db)
	players,err := g.GetPlayers(context.TODO())
	require.Nil(t, err) // require is required to address custom error which implements Error interface
	assert.NotNil(t, players)

}

//GetPlayer test
func TestGetPlayer (t *testing.T) {
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockPlayers := []models.Player{
			models.Player {
				PlayerId: 1001,
				NamePlayer: "Tom",
				ChipsAmount : 200,
				TotalNoOfChips: 2,
			},
	}
	rows := sqlmock.NewRows([]string{"idPlayers","namePlayer","chipsAmount","totalNoChips"}).
			AddRow(mockPlayers[0].PlayerId,mockPlayers[0].NamePlayer,mockPlayers[0].ChipsAmount,mockPlayers[0].TotalNoOfChips)
	query := "SELECT idPlayers,namePlayer,chipsAmount,totalNoChips from Players where idPlayers=\\?"	
	mock.ExpectQuery(query).WillReturnRows(rows)
	p := repo.NewPlayerRepository(db)
	player,err := p.GetPlayer(context.TODO(),1002)
	require.Nil(t, err) // require is required to address custom error which implements Error interface
	assert.NotNil(t, player)	

}


//CreatePlayer test
func TestCreatePlayer(t *testing.T) {
	py := models.Player{
		NamePlayer: "Tom",
		ChipsAmount : 200,
		TotalNoOfChips: 2,	
	}	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}	
	defer db.Close()
	query := "insert into Players\\(namePlayer,chipsAmount,totalNoChips\\) values\\(\\?,\\?,\\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(py.NamePlayer,py.ChipsAmount,py.TotalNoOfChips).WillReturnResult(sqlmock.NewResult(1005, 1))
	p := repo.NewPlayerRepository(db)
	flag,e := p.CreatePlayer(context.TODO(), py)
	require.Nil(t, e)
	assert.Equal(t, flag, true)
}