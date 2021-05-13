package log_controller
import(
	
    "os"
    "path/filepath"
)
type Loglist struct {
	UserName string   `json:"userName"`
	Project  string   `json:"project"`
	Logs     []string `json:"logs"`
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
