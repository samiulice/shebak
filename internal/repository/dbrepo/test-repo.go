package dbrepo

import (
	"github.com/samiulru/shebak/internal/models"
)


// GetAllServiceCategoryMain returns a slice of existant Service Category Main from the database
func (m *testDBRepo) GetAllServiceCategoryMain() ([]models.ServiceCategoryMain, error) {
	return nil, nil
}

// GetServiceCategoryMainByID searches and returns Service Category Main info by ID from the database
func (m *testDBRepo) GetServiceCategoryMainByID(id int) (models.ServiceCategoryMain, error){

	var category models.ServiceCategoryMain
	return category, nil
}

// InsertServiceCategoryMain inserts CategoryMain info to the database
func (m *testDBRepo) InsertServiceCategoryMain(service models.ServiceCategoryMain) (int, error) {
	return 0, nil
}

// InsertServiceCategorySub inserts CategorySub info to the database
func (m *testDBRepo) InsertServiceCategorySub(service models.ServiceCategorySub) (int, error) {
	return 0, nil
}

// GetAllServiceCategorySub returns a slice of all Sub Categories from the database
func (m *testDBRepo) GetAllServiceCategorySub() ([]models.ServiceCategorySub, error) {
	var allSubCategory []models.ServiceCategorySub
	return allSubCategory, nil
}
// GetServiceCategorySubByID searches and returns Service Category Sub info by ID from the database
func (m *testDBRepo) GetServiceCategorySubByID(id int) (models.ServiceCategorySub, error) {
	var categorySub models.ServiceCategorySub
	return categorySub, nil
}

//GetSubListByMainID returns a slice of Service Category Sub based on foreign key category_id
func (m *testDBRepo) GetSubListByMainID(category_id int) ([]models.ServiceCategorySub, error){
	var list []models.ServiceCategorySub
	return list, nil
}

// GetAllServices returns a slice of all Services from the database
func (m *testDBRepo) GetAllServices() ([]models.Service, error) {
	var allServices []models.Service
	return allServices, nil
}
//GetServiceListByMainID_SubID returns a slice of Services based on foreign keys category_id and sub_category_id
func (m *testDBRepo) GetServiceListByMainID_SubID(category_id, sub_category_id int) ([]models.Service, error){
	var service []models.Service
	return service, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////Deprecated:: Old functions needs to be updated////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////// 

//AllUsers retrive all users from database
func (m *testDBRepo) AllAdmins() bool {
	return true
}

// GetUserByID searches user by ID
func (m *testDBRepo) GetAdminByID(id int) (models.Admin, error) {
	var u models.Admin
	return u, nil
}

// UpdateUser updates a user in the database
func (m *testDBRepo) UpdateAdmin(u models.Admin) error {
	return nil
}

// Authenticate authenticates a user
func (m *testDBRepo) Authenticate(email, testPassword string) (int,  string, int, error) {
	return 0, "", 0, nil
}