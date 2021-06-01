package repository

import(
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
)

type LogRepository interface {
	Save(log datamodels.Log)
	GetAll()
	GetByUserName(user string)
	GetByProjectname(projectname string)
}
