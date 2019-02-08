package models

type Game struct {
	
	GameId int `json:"IdGame"`
	GameName string `json:"Name"`
	QuestOrder string `json:"QuestOrder"`
}