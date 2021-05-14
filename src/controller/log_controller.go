package log_controller

import (
	//"io/ioutil"
	"io/ioutil"
	"os"
	"path/filepath"
	//"io/ioutil"
)
type Loglist struct {
	UserName string   `json:"userName"`
	Project  string   `json:"project"`
	Logs     []string `json:"logs"`
}

type LogContent struct {

    FileName string `json:"filename"`
	Content string  `json:"content"`
	
}



func GetFileList(user string, project string) Loglist {
	
	
 
var files []string
   // user :="tharindu"
    //project := "project1"
    root := "logs/"+user+"/"+project
    //root:= "../logs/" + user + "/" + project
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, info.Name())
        return nil
    })
    if err != nil {
        panic(err)
    }
    

	loglist := Loglist{
		UserName: user,
		Project:  project,
		Logs:     files,
	}

   
	return loglist
}

func GetLogfileContent(user string, project string ,Logs string) LogContent{

   
   root := "logs/"+user+"/"+project
   
    data, err := ioutil.ReadFile(root+"/"+Logs + ".txt")
	if err != nil {
	panic(err)
	 }

	var dataT = string(data)

	
     logcontent := LogContent{
            FileName: Logs,
            Content: dataT,

     }

     return logcontent

	
}
