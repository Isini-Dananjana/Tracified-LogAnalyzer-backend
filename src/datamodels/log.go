package datamodels



type Log struct {
	Username    string `json:"username"`
	ProjectName string `json:"projectname"`
	LogFileName string `json:"logfilename"`
	LastUpdate  string `json:"lastupdate"`
	FileId  		string `json:"fileId"`
}
