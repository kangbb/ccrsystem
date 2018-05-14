package services

import (
	"github.com/kangbb/ccrsystem/models/entities"
)

type OrganizationInfoService struct{}

var OrganizationService = OrganizationInfoService{}

/*
* Create a new organization
 */
func (*OrganizationInfoService) NewOrganization(name string, introduction string) *entities.OrganizationInfo {
	org := &entities.OrganizationInfo{
		OrganizationName: name,
		Introduction:     introduction,
	}
	return org
}

/*
* Insert a organization information to the db
 */
func (*OrganizationInfoService) SaveAInfo(org *entities.OrganizationInfo) error {
	_, err := entities.MasterEngine.InsertOne(org)

	return err
}

/*
* Get a organization information by ID
 */
func (*OrganizationInfoService) FindInfoById(id int) (*entities.OrganizationInfo, error) {
	org := new(entities.OrganizationInfo)
	_, err := entities.SlaveEngine.Id(id).Get(org)

	return org, err
}

/*
* Get a organization information by note
 */
func (*OrganizationInfoService) FindInfoByNote(note string) (*entities.OrganizationInfo, error) {
	org := new(entities.OrganizationInfo)
	_, err := entities.SlaveEngine.Where("note = ?", note).Get(org)

	return org, err
}

/*
* Get a organization information by order
 */
func (*OrganizationInfoService) FindInfoByOrder(order int) (*entities.OrganizationInfo, error) {
	org := new(entities.OrganizationInfo)
	_, err := entities.SlaveEngine.Where("order = ?", order).Get(org)

	return org, err
}

/*
* Find all organizations information
 */
func (*OrganizationInfoService) FindAllInfo() ([]entities.OrganizationInfo, error) {
	orgs := make([]entities.OrganizationInfo, 0)
	err := entities.SlaveEngine.Asc("order").Find(&orgs)

	return orgs, err
}

/*
* Update organization information
 */
func (*OrganizationInfoService) UpdateInfo(id int, name string, introduction string) error {
	org := new(entities.OrganizationInfo)
	org.OrganizationId = id
	org.OrganizationName = name
	org.Introduction = introduction

	_, err := entities.MasterEngine.Id(id).Update(org)

	return err
}

/*
* Delete a piece of organization information
 */
func (*OrganizationInfoService) DeleteInfo(id int) error {
	org := new(entities.OrganizationInfo)
	_, err := entities.MasterEngine.Id(id).Delete(org)

	return err
}
