package services

import (
	"github.com/kangbb/ccrsystem/models/entities"
)

type AdminInfoService struct{}

var AdminService = AdminInfoService{}

/*
* Create a new Admin
 */
func (*AdminInfoService) NewAdmin(id int, pwd string, name string) *entities.AdminInfo {
	admin := &entities.AdminInfo{
		AdminId:   id,
		AdminPwd:  pwd,
		AdminName: name,
	}
	return admin
}

/*
* Insert admin information to the db
 */
func (*AdminInfoService) SaveAInfo(adm *entities.AdminInfo) error {
	_, err := entities.MasterEngine.InsertOne(adm)
	return err
}

/*
* Get a admin information by ID
 */
func (*AdminInfoService) FindInfoById(id int) (*entities.AdminInfo, error) {
	adm := new(entities.AdminInfo)
	_, err := entities.SlaveEngine.Id(id).Get(adm)
	return adm, err
}

/*
* Find all admin information
 */
func (*AdminInfoService) FindAllInfo() ([]entities.AdminInfo, error) {
	adm := make([]entities.AdminInfo, 0)
	err := entities.SlaveEngine.Find(&adm)

	return adm, err
}

/*
* Update admin information, just for password.
 */
func (*AdminInfoService) UpdateInfo(id int, pwd string) error {
	adm := new(entities.AdminInfo)
	adm.AdminPwd = pwd

	_, err := entities.MasterEngine.Id(id).Update(adm)

	return err
}

/*
* Delete a piece of admin information
 */
func (*AdminInfoService) DeleteInfo(id int) error {
	adm := new(entities.AdminInfo)
	_, err := entities.MasterEngine.Id(id).Delete(adm)

	return err
}
