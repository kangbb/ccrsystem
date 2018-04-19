package service

import (
	"github.com/kangbb/class-reservation/core/models/entities"
)

type AdminInfoService struct{}

var AdminService = AdminInfoService{}

//Create a new Admin
func (*AdminInfoService) NewAdmin(id int, pwd string, name string, perm bool) entities.AdminInfo {
	admin := entities.AdminInfo{
		AdminId:    id,
		AdminPwd:   pwd,
		AdminName:  name,
		Permission: perm,
	}
	return admin
}

//Insert admin information to the db
func (*AdminInfoService) SaveAInfo(adm *entities.AdminInfo) (bool, error) {
	_, err := entities.MasterEngine.InsertOne(adm)
	if err != nil {
		panic(err)
	}
	return true, nil
}

//Get a admin information by ID
func (*AdminInfoService) FindInfoById(id int) *entities.AdminInfo {
	adm := new(entities.AdminInfo)
	_, err := entities.SlaveEngine.Id(id).Get(adm)
	if err != nil {
		panic(err)
	}
	return adm
}

//Find all admin information
func (*AdminInfoService) FindAllInfo() []entities.AdminInfo {
	adm := make([]entities.AdminInfo, 0)
	err := entities.SlaveEngine.Find(&adm)
	if err != nil {
		panic(err)
	}
	return adm
}

//Update admin information
//Just for password, permission
func (*AdminInfoService) UpdateInfo(id int, pwd string, arg ...interface{}) (bool, error) {
	adm := new(entities.AdminInfo)
	switch len(arg) {
	case 1:
		adm.Permission = arg[0].(bool)
	case 0:
		adm.AdminPwd = pwd
	}
	_, err := entities.MasterEngine.Id(id).Update(adm)
	if err != nil {
		panic(err)
	}
	return true, nil
}

//Delete admin information
func (*AdminInfoService) DeleteInfo(id int) (bool, error) {
	adm := new(entities.AdminInfo)
	_, err := entities.MasterEngine.Id(id).Delete(adm)
	if err != nil {
		panic(err)
	}
	return false, nil
}
