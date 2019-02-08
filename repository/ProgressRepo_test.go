package repository_test

import (
	repo "github.com/siddhiparekh11/GoChallenge/repository"
	"github.com/siddhiparekh11/GoChallenge/models"
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"context"
	"time"
)

/*
	Test file for Progress Repo. Mocking the database and testing against predefined input. 
*/


//GetPlayerState test 
func TestGetPlayerState (t *testing.T) {
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockProgress := []models.Progress{
			models.Progress {
				GameId: 101,
				QuestId: 11, 
				PlayerId: 1001,  
				QuestPointsEarned: 203,
				TotQuestCompPer: 100,
				LastMilestoneIndex: 1,
				LastMilestone: models.CusMilestone{1,30},
				Milestones: []models.CusMilestone{{1,30},{2,60}},
				CreatedTimestamp: time.Now(),
				ChipAmountBet: 200,
			},
	}

	rows := sqlmock.NewRows([]string{"idGame","idQuest","idPlayer","questPointsEarned","totalQuestComPercent","lastMilestoneInd","createdTimestamp"}).
			AddRow(mockProgress[0].GameId,mockProgress[0].QuestId,mockProgress[0].PlayerId,mockProgress[0].QuestPointsEarned,mockProgress[0].TotQuestCompPer,mockProgress[0].LastMilestoneIndex,mockProgress[0].CreatedTimestamp)
	query := "SELECT idGame,idQuest,idPlayer,questPointsEarned,totalQuestComPercent,lastMilestoneInd,createdTimestamp FROM Progress where idPlayer=\\? order by createdTimestamp desc limit 1"	
	mock.ExpectQuery(query).WillReturnRows(rows)
	pr := repo.NewProgressRepository(db)
	pstate,err := pr.GetPlayerState(context.TODO(),11)
	require.Nil(t, err) // require is required to address custom error which implements Error interface
	assert.NotNil(t, pstate)
}

//GetAllQuestProgress test
func TestGetAllQuestProgress(t *testing.T){
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockProgress := []models.Progress{
			models.Progress {
				GameId: 101,
				QuestId: 11, 
				PlayerId: 1001,  
				QuestPointsEarned: 203,
				TotQuestCompPer: 100,
				LastMilestoneIndex: 1,
				LastMilestone: models.CusMilestone{1,30},
				Milestones: []models.CusMilestone{{1,30},{2,60}},
				CreatedTimestamp: time.Now(),
				ChipAmountBet: 200,
			},
			models.Progress {
				GameId: 101,
				QuestId: 12, 
				PlayerId: 1001,  
				QuestPointsEarned: 103,
				TotQuestCompPer: 50,
				LastMilestoneIndex: 1,
				LastMilestone: models.CusMilestone{1,30},
				Milestones: []models.CusMilestone{{1,30},{2,60}},
				CreatedTimestamp: time.Now(),
				ChipAmountBet: 300,
			},
	}

	rows := sqlmock.NewRows([]string{"idGame","idQuest","idPlayer","questPointsEarned","totalQuestComPercent","lastMilestoneInd","createdTimestamp"}).
			AddRow(mockProgress[0].GameId,mockProgress[0].QuestId,mockProgress[0].PlayerId,mockProgress[0].QuestPointsEarned,mockProgress[0].TotQuestCompPer,mockProgress[0].LastMilestoneIndex,mockProgress[0].CreatedTimestamp)
	query := "SELECT idGame,idQuest,idPlayer,questPointsEarned,totalQuestComPercent,lastMilestoneInd,createdTimestamp FROM Progress where idPlayer=\\? and idGame=\\? order by createdTimestamp"	
	mock.ExpectQuery(query).WillReturnRows(rows)
	pr := repo.NewProgressRepository(db)
	pstate,err := pr.GetAllQuestProgress(context.TODO(),1001,101)
	require.Nil(t, err) // require is required to address custom error which implements Error interface
	assert.NotNil(t, pstate)
}


//UpdateProgress test
func TestUpdateProgress(t *testing.T) {
	tm := time.Now()
	pr := models.Progress {
				GameId: 101,
				QuestId: 11, 
				PlayerId: 1001,  
				QuestPointsEarned: 203,
				TotQuestCompPer: 100,
				LastMilestoneIndex: 1,
				LastMilestone: models.CusMilestone{1,30},
				Milestones: []models.CusMilestone{{1,30},{2,60}},
				CreatedTimestamp: tm,
				ChipAmountBet: 200,
			}				
	ans := &models.Progress {
				GameId: 101,
				QuestId: 11, 
				PlayerId: 1001,  
				QuestPointsEarned: 203,
				TotQuestCompPer: 100,
				LastMilestoneIndex: 1,
				LastMilestone: models.CusMilestone{1,30},
				Milestones: []models.CusMilestone{{1,30},{2,60}},
				CreatedTimestamp: tm,
				ChipAmountBet: 200,
			}	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}	
	defer db.Close()
	query := "insert into Progress\\(idGame,idQuest,idPlayer,questPointsEarned,totalQuestComPercent,lastMilestoneInd,createdTimestamp\\) values\\(\\?,\\?,\\?,\\?,\\?,\\?,\\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(pr.GameId,pr.QuestId,pr.PlayerId,pr.QuestPointsEarned,pr.TotQuestCompPer,pr.LastMilestoneIndex,pr.CreatedTimestamp).WillReturnResult(sqlmock.NewResult(3, 1))
	prg := repo.NewProgressRepository(db)
	p,e := prg.UpdateProgress(context.TODO(), pr)
	require.Nil(t, e)
	assert.Equal(t, p,ans)
}