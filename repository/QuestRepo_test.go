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
	Test file for Quest Repo. Mocking the database and testing against predefined input. 
*/


//GetQuests test
func TestGetQuests (t *testing.T) {
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockQuests := []models.Quest{
			models.Quest {
				QuestId: 11,
				QuestName: "Level4",
				MilestonesOrder:"2,3",
			},
			models.Quest {
				QuestId: 12,
				QuestName: "Level5",
				MilestonesOrder:"4,5",
			},
	}
	rows := sqlmock.NewRows([]string{"idQuests","questName","milestonesOrder"}).
			AddRow(mockQuests[0].QuestId,mockQuests[0].QuestName,mockQuests[0].MilestonesOrder).
			AddRow(mockQuests[1].QuestId,mockQuests[1].QuestName,mockQuests[1].MilestonesOrder)
	query := "SELECT idQuests,questName,milestonesOrder from Quests"	
	mock.ExpectQuery(query).WillReturnRows(rows)
	q := repo.NewQuestRepository(db)
	quests,err := q.GetQuests(context.TODO())
	require.Nil(t, err) // require is required to address custom error which implements Error interface
	assert.NotNil(t, quests)
}

//GetQuest test
func TestGetQuest (t *testing.T) {
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockQuests := []models.Quest{

			models.Quest {
				QuestId: 14,
				QuestName: "Level4",
				MilestonesOrder:"2,3",
			},
	}
	rows := sqlmock.NewRows([]string{"idQuests","questName","milestonesOrder"}).
			AddRow(mockQuests[0].QuestId,mockQuests[0].QuestName,mockQuests[0].MilestonesOrder)
	query := "SELECT idQuests,questName,milestonesOrder from Quests where idQuests=\\?"	
	mock.ExpectQuery(query).WillReturnRows(rows)
	q := repo.NewQuestRepository(db)
	quest,err := q.GetQuest(context.TODO(),14)
	require.Nil(t, err) // require is required to address custom error which implements Error interface
	assert.NotNil(t, quest)
}


//CreateQuest test
func TestCreateQuest(t *testing.T) {
	qr := models.Quest {
				QuestName: "Level5",
				MilestonesOrder:"1,2,3",
			}	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}	
	defer db.Close()
	query := "insert into Quests\\(questName,milestonesOrder\\) values\\(\\?,\\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(qr.QuestName,qr.MilestonesOrder).WillReturnResult(sqlmock.NewResult(15, 1))
	q := repo.NewQuestRepository(db)
	flag,e := q.CreateQuest(context.TODO(), qr)
	require.Nil(t, e)
	assert.Equal(t, flag, true)
}