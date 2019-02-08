package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"context"
	"database/sql"
	"github.com/siddhiparekh11/GoChallenge/interfaces"
	"github.com/siddhiparekh11/GoChallenge/models"
)


type QuestRepository struct {
    Conn *sql.DB
}


func NewQuestRepository(conn *sql.DB) (interfaces.IQuest) {
		return &QuestRepository {conn}
}


func (questRepo *QuestRepository) CreateQuest(ctx context.Context, quest models.Quest) (bool,*models.QError) {
	insert, err := questRepo.Conn.Prepare("insert into Quests(questName,milestonesOrder) values(?,?)")
	if err != nil {
		return false, &models.QError{Where: "QuestRepo/CreateQuest", What: err.Error()}
	}	
	_, err = insert.Exec(quest.QuestName,quest.MilestonesOrder)
	if err != nil {
		return false,&models.QError{Where: "QuestRepo/CreateQuest", What: err.Error()}
	}
	return true,nil
}

func (questRepo *QuestRepository) GetQuests(ctx context.Context) ([] *models.Quest, *models.QError) {
	query := "SELECT idQuests,questName,milestonesOrder from Quests"
	rows,err := questRepo.Conn.QueryContext(ctx,query)
	if err!=nil {
		return nil,&models.QError{Where: "QuestRepo/GetQuests", What: err.Error()}
	}
	defer rows.Close()
	quests := make([] *models.Quest,0)
	for rows.Next() {		
		q := new(models.Quest)
		err = rows.Scan(&q.QuestId,&q.QuestName,&q.MilestonesOrder)		
		if err!=nil {
		return nil,&models.QError{Where: "QuestRepo/GetQuests", What: err.Error()}
		}
		quests = append(quests,q)	
	}
	if len(quests) == 0 {
		return nil, &models.QError{Where: "QuestRepo/GetGames", What: "There are no existing quests"}
	}
	return quests,nil
}

func (questRepo *QuestRepository) GetQuest(ctx context.Context,questId int) (*models.Quest, *models.QError) {
	query := "SELECT idQuests,questName,milestonesOrder from Quests where idQuests=?" 
	rows,err := questRepo.Conn.QueryContext(ctx,query,questId)
	if err!=nil {		
		return nil,&models.QError{Where: "QuestRepo/GetQuest", What: err.Error()}
	}
	defer rows.Close()
	quests := make([] *models.Quest,0)
	for rows.Next() {		
		q := new(models.Quest)
		err = rows.Scan(&q.QuestId,&q.QuestName,&q.MilestonesOrder)		
		if err!=nil {
		return nil,&models.QError{Where: "QuestRepo/GetQuest", What: err.Error()}
		}
		quests = append(quests,q)	
	}
	if len(quests) == 0 {
		return nil, &models.QError{Where: "QuestRepo/GetQuest", What: "No such Quest Id"}
	}
	return quests[0],nil

}
