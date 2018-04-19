package service

type InfoService interface {
	SaveAInfo(Info interface{}) (bool, error)
	FindInfoById(id int) interface{}
	FindAllInfo() interface{}
	UpdateInfo(id int, arg ...interface{}) (bool, error)
	DeleteInfo(id int, arg ...interface{}) (bool, error)
}
