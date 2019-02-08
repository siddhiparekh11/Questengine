package models

type Quest struct {

	QuestId int `json:"IdQuest"`
	QuestName string `json:"Name"`
	MilestonesOrder string `json:"MilestonesOrder"`
}