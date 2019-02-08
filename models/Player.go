package models

type Player struct {

	PlayerId int `json:"IdPlayer"`
	NamePlayer string `json:"Name"`
	ChipsAmount int `json:"ChipsAmount"`
	TotalNoOfChips int `json:"TotalNoChips"`
}