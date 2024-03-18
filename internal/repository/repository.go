package repository

import (
	"github.com/samiulru/shebak/internal/models"
)

type DatabaseRepo interface {

	//New--Start

	GetAllServiceCategoryMain() ([]models.ServiceCategoryMain, error)
	GetServiceCategoryMainByID(id int) (models.ServiceCategoryMain, error)
	InsertServiceCategoryMain(service models.ServiceCategoryMain) (int, error)

	InsertServiceCategorySub(service models.ServiceCategorySub) (int, error)

	GetAllServiceCategorySub() ([]models.ServiceCategorySub, error)
	GetServiceCategorySubByID(id int) (models.ServiceCategorySub, error)
	GetSubListByMainID(category_id int) ([]models.ServiceCategorySub, error)

	GetAllServices() ([]models.Service, error)
	GetServiceListByMainID_SubID(category_id, sub_category_id int) ([]models.Service, error)
	//New--End

	//OLD
	AllAdmins() bool
	GetAdminByID(id int) (models.Admin, error)
	UpdateAdmin(u models.Admin) error
	Authenticate(email, testPassword string) (int, string, int, error)
}
