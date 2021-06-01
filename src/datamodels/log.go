package datamodels

import "time"

type Log struct {
	Username    string `json:"uname"`
	ProjectName string `json:"proname"`
	LogFileName string `json:"log"`
	LastUpdate  time.Time `json:"lastUP"`
}
