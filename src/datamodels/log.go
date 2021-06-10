package datamodels



type Log struct {
	Username    string `json:"uname"`
	ProjectName string `json:"proname"`
	LogFileName string `json:"log"`
	LastUpdate  string `json:"lastUP"`
}
