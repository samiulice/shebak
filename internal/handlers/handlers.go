package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/samiulru/shebak/internal/config"
	"github.com/samiulru/shebak/internal/driver"
	"github.com/samiulru/shebak/internal/forms"
	"github.com/samiulru/shebak/internal/models"
	"github.com/samiulru/shebak/internal/render"
	"github.com/samiulru/shebak/internal/repository"
	"github.com/samiulru/shebak/internal/repository/dbrepo"
)

// Repo is the Repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewTestRepo creates a testing repository
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingRepo(a),
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// ServiceList returns a slice of current services from then database
func (m *Repository) CategoryMainList() []models.ServiceCategoryMain {
	CategoryMain, err := m.DB.GetAllServiceCategoryMain()
	if err != nil {
		panic(err)
	}
	return CategoryMain
}

// ServiceList returns a slice of current services from then database
func (m *Repository) CategorySubList() []models.ServiceCategorySub {
	CategorySub, err := m.DB.GetAllServiceCategorySub()
	if err != nil {
		panic(err)
	}
	return CategorySub
}

// ServiceList returns a slice of current services from then database
func (m *Repository) ServiceList() []models.Service {
	serviceList, err := m.DB.GetAllServices()
	if err != nil {
		panic(err)
	}
	return serviceList
}

////////////////////////////////Public Handler///////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////

// Home handles the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	_ = render.TemplatesRenderer(w, r, "home.page.tmpl", &models.TemplateData{})
}

/*......................................................................
....................Admin Tools Handler Functions....................
......................................................................*/

// UserLogin handles UserLogin page
func (m *Repository) AdminLogin(w http.ResponseWriter, r *http.Request) {
	_ = render.TemplatesRenderer(w, r, "login.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostUserLogin handles authentication and Login of users
func (m *Repository) PostAdminLogin(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")

	if !form.Valid() {
		log.Println(err)
		render.TemplatesRenderer(w, r, "login.page.tmpl", &models.TemplateData{
			Form: form,
		})
		http.Error(w, "", http.StatusSeeOther)
		return
	}

	id, _, accessLevel, err := m.DB.Authenticate(email, password)
	if err != nil {
		log.Println(err)
		m.App.Session.Put(r.Context(), "error", "invalid login credentials")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}
	m.App.Session.Put(r.Context(), "user_id", id)
	m.App.Session.Put(r.Context(), "user_access_level", accessLevel)
	
	// AccessLevel := 1 means SuperAdmin
	// AccessLevel := 2 means Admin
	// AccessLevel := 3 means Employee
	// AccessLevel := 4 means User
	url := []string{"", "/private/admin/super/dashboard", "/private/admin/dashboard", "/employee/dashboard", "/user/dashboard"}
	http.Redirect(w, r, url[accessLevel], http.StatusSeeOther)
}

// AdminLogout logs an Admin out
func (m *Repository) AdminLogout(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}

// //////////////////////////////Super Admin Handler///////////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////////////////
// AdminDashboard handles Admins dashboard
func (m *Repository) SuperAdminDashboard(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	render.TemplatesRenderer(w, r, "super-admin-dashboard.page.tmpl", &models.TemplateData{
		Data:         data,
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
	})
}

////////////////////////////////Admin Handler///////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////

// AdminDashboard handles Admins dashboard
func (m *Repository) AdminDashboard(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	render.TemplatesRenderer(w, r, "admin-dashboard.page.tmpl", &models.TemplateData{
		Data:         data,
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
	})
}

// AdminSubItemList renders sub item info
func (m *Repository) AdminSubItemList(w http.ResponseWriter, r *http.Request) {
	uriParts := strings.Split(r.RequestURI, "/")
	mainID, _ := strconv.Atoi(uriParts[4]) //4th element : mainID

	currentMain, err := m.DB.GetServiceCategoryMainByID(mainID)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Failed to get category info from database")
		http.Redirect(w, r, "/private/admin/dashboard", http.StatusSeeOther)
		return

	}
	subList, err := m.DB.GetSubListByMainID(mainID)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", err)
		http.Redirect(w, r, "/private/admin/dashboard", http.StatusSeeOther)
		return
	}
	data := make(map[string]interface{})
	data["current_main"] = currentMain
	data["sub_list"] = subList

	if strings.Contains(r.RequestURI, "sub") {
		subID, _ := strconv.Atoi(uriParts[5]) //5th element : subID (since it does exist)
		currentSub, err := m.DB.GetServiceCategorySubByID(subID)
		if err != nil {
			m.App.Session.Put(r.Context(), "error", "Failed to get sub category info from database")
			http.Redirect(w, r, "/private/admin/dashboard", http.StatusSeeOther)
			return
		}
		serviceList, err := m.DB.GetServiceListByMainID_SubID(mainID, subID)
		if err != nil {
			m.App.Session.Put(r.Context(), "error", "Failed to get service list from database")
			http.Redirect(w, r, "/private/admin/dashboard", http.StatusSeeOther)
			return
		}
		data["current_sub"] = currentSub
		data["service_list"] = serviceList
		data["sub"] = "sub"
	}
	render.TemplatesRenderer(w, r, "admin-sub-item-list.page.tmpl", &models.TemplateData{
		Data:         data,
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
	})
}

// AdminAddNewCategory renders adding new category page
func (m *Repository) AdminAddNewCategory(w http.ResponseWriter, r *http.Request) {

	//////////////////////////////////////////////////////////////////////
	/////////////////////////////////TODO/////////////////////////////////
	//////////////////////////////////////////////////////////////////////
	// var mainCategory models.ServiceCategoryMain
	data := make(map[string]interface{})
	render.TemplatesRenderer(w, r, "admin-add-new-category.page.tmpl", &models.TemplateData{
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
		Data:         data,
		Form:         forms.New(nil),
	})
}

// PostAdminAddNewCategory handles post request for adding new category
func (m *Repository) PostAdminAddNewCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Failed to parse form! Try agian")
		http.Redirect(w, r, "/private/admin/category/add-new", http.StatusSeeOther)
		return
	}
	// Get image from the form
	file, header, err := r.FormFile("category_thumbnail")
	if err != nil && err != http.ErrMissingFile {
		m.App.Session.Put(r.Context(), "error", "Unable to read image")
		http.Redirect(w, r, "/private/admin/category/add-new", http.StatusSeeOther)
		return
	}
	if err != http.ErrMissingFile {
		defer file.Close()
	}

	// Getting Other informations
	var category models.ServiceCategoryMain
	category.Name = r.Form.Get("category_name")
	category.Description = r.Form.Get("category_description")
	category.Available = 1

	form := forms.New(r.PostForm)
	form.Required("category_name", "category_description")
	form.MinLength("category_name", 1)
	form.MinLength("category_description", 10)
	if err == http.ErrMissingFile {
		form.AddErr("category_thumbnail")
	}

	if !form.Valid() { //Invalid user input
		log.Println("Invalid form")
		data := make(map[string]interface{})
		data["current_main_category"] = category
		render.TemplatesRenderer(w, r, "admin-add-new-category.page.tmpl", &models.TemplateData{
			Data:         data,
			Form:         form,
			CategoryMain: m.App.CategoryMain,
			CategorySub:  m.App.CategorySub,
			Services:     m.App.Services,
		})
		return
	}

	//Give an arbitary name for the image
	nameSplit := strings.Split(header.Filename, ".")
	category.Thumbnanil = strings.ReplaceAll(category.Name, " ", "") + "_" + uuid.New().String() + "." + nameSplit[1]

	// Create a new file in the server's filesystem
	dst, err := os.Create(filepath.Join("./static/public/images/main-categories-thumbnail", category.Thumbnanil))
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Failed to create image file in the server's filesystem! Try agian")
		http.Redirect(w, r, "/private/admin/category/add-new", http.StatusSeeOther)
		return
	}
	defer dst.Close()

	//Copy the uploaded file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Failed to copy image to the destination! Try agian")
		http.Redirect(w, r, "/private/admin/category/add-new", http.StatusSeeOther)
		return
	}

	_, err = m.DB.InsertServiceCategoryMain(category)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Unable to insert data to the database! Try agian")
		http.Redirect(w, r, "/private/admin/category/add-new", http.StatusSeeOther)
		return
	}

	m.App.Session.Put(r.Context(), "flash", "Service Category info inserted successfully")

	//Updating servicelist automatically
	categoryMain, err := m.DB.GetAllServiceCategoryMain() //make an manual update func for this update process
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Database error, Serviclist needs to be updated manually")
		http.Redirect(w, r, "/private/admin/category/add-new", http.StatusSeeOther)
		return
	}
	m.App.CategoryMain = categoryMain
	http.Redirect(w, r, "/private/admin/category/add-new", http.StatusSeeOther)

}

// AdminUpdateCategory renders update existing category page
func (m *Repository) AdminUpdateCategory(w http.ResponseWriter, r *http.Request) {
	//////////////////////////////////////////////////////////////////////
	/////////////////////////////////TODO/////////////////////////////////
	//////////////////////////////////////////////////////////////////////
	data := make(map[string]interface{})

	render.TemplatesRenderer(w, r, "admin-update-category.page.tmpl", &models.TemplateData{
		Data:         data,
		Form:         forms.New(nil),
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
	})
}

// PostAdminUpdateCategory handles post request to update existing category
func (m *Repository) PostAdminUpdateCategory(w http.ResponseWriter, r *http.Request) {
	//////////////////////////////////////////////////////////////////////
	/////////////////////////////////TODO/////////////////////////////////
	//////////////////////////////////////////////////////////////////////
	m.App.Session.Put(r.Context(), "success", "Category updated success")
	http.Redirect(w, r, "/private/admin/dashboard", http.StatusSeeOther)
}

// AdminAddNewSubCategory renders adding new sub category page
func (m *Repository) AdminAddNewSubCategory(w http.ResponseWriter, r *http.Request) {

	//////////////////////////////////////////////////////////////////////
	/////////////////////////////////TODO/////////////////////////////////
	//////////////////////////////////////////////////////////////////////
	data := make(map[string]interface{})

	render.TemplatesRenderer(w, r, "admin-add-new-sub-category.page.tmpl", &models.TemplateData{
		Data:         data,
		Form:         forms.New(nil),
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
	})
}

// PostAdminAddNewSubCategory handles post request for adding new sub category
func (m *Repository) PostAdminAddNewSubCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Failed to parse form! Try agian")
		http.Redirect(w, r, "/private/admin/sub-category/add-new", http.StatusSeeOther)
		return
	}
	// Get image from the form
	file, header, err := r.FormFile("sub_category_thumbnail")
	if err != nil && err != http.ErrMissingFile {
		m.App.Session.Put(r.Context(), "error", "Unable to read image")
		http.Redirect(w, r, "/private/admin/sub-category/add-new", http.StatusSeeOther)
		return
	}
	if err != http.ErrMissingFile {
		defer file.Close()
	}

	// Getting Other informations
	var subCategory models.ServiceCategorySub
	subCategory.Name = r.Form.Get("sub_category_name")
	subCategory.Description = r.Form.Get("sub_category_description")
	subCategory.Available = 1
	catID, _ := strconv.Atoi(r.Form.Get("seleced_category"))
	subCategory.CategoryID = catID

	form := forms.New(r.PostForm)
	form.Required("seleced_category", "sub_category_name", "sub_category_description")
	form.MinLength("sub_category_name", 1)
	form.MinLength("sub_category_description", 10)
	if err == http.ErrMissingFile {
		form.IsImage("sub_category_thumbnail")
	}
	if catID == 0 {
		form.AddErr("seleced_category")
	}

	if !form.Valid() { //Invalid user input
		log.Println("Invalid form")
		data := make(map[string]interface{})
		data["current_sub_category"] = subCategory
		render.TemplatesRenderer(w, r, "admin-add-new-sub-category.page.tmpl", &models.TemplateData{
			Data:         data,
			Form:         form,
			CategoryMain: m.App.CategoryMain,
			CategorySub:  m.App.CategorySub,
			Services:     m.App.Services,
		})
		return
	}

	//Give an arbitary name for the image
	nameSplit := strings.Split(header.Filename, ".")
	subCategory.Thumbnanil = strings.ReplaceAll(subCategory.Name, " ", "") + "_" + uuid.New().String() + "." + nameSplit[1]

	// Create a new file in the server's filesystem
	dst, err := os.Create(filepath.Join("./static/public/images/sub-categories-thumbnail", subCategory.Thumbnanil))
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Failed to create image file in the server's filesystem! Try agian")
		http.Redirect(w, r, "/private/admin/sub-category/add-new", http.StatusSeeOther)
		return
	}
	defer dst.Close()

	//Copy the uploaded file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		os.Remove(filepath.Join("./static/public/images/sub-categories-thumbnail", subCategory.Thumbnanil))
		m.App.Session.Put(r.Context(), "error", "Failed to copy image to the destination! Try agian")
		http.Redirect(w, r, "/private/admin/sub-category/add-new", http.StatusSeeOther)
		return
	}

	_, err = m.DB.InsertServiceCategorySub(subCategory)
	if err != nil {
		os.Remove(filepath.Join("./static/public/images/sub-categories-thumbnail", subCategory.Thumbnanil))
		m.App.Session.Put(r.Context(), "error", "Unable to insert data to the database! Try agian")
		http.Redirect(w, r, "/private/admin/sub-category/add-new", http.StatusSeeOther)
		return
	}

	m.App.Session.Put(r.Context(), "flash", "Service Category info inserted successfully")
	//Updating servicelist automatically
	categorySub, err := m.DB.GetAllServiceCategorySub() //make an manual update func for this update process
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Database error, Serviclist needs to be updated manually")
		http.Redirect(w, r, "/private/admin/sub-category/add-new", http.StatusSeeOther)
		return
	}
	m.App.CategorySub = categorySub
	http.Redirect(w, r, "/private/admin/sub-category/add-new", http.StatusSeeOther)

}

// AdminAddNewSubCategory renders update existing sub category page
func (m *Repository) AdminUpdateSubCategory(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	render.TemplatesRenderer(w, r, "admin-update-sub-category.page.tmpl", &models.TemplateData{
		Data:         data,
		Form:         forms.New(nil),
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
	})
}

// PostAdminUpdateSubCategory handles post request to update existing sub category
func (m *Repository) PostAdminUpdateSubCategory(w http.ResponseWriter, r *http.Request) {

	//////////////////////////////////////////////////////////////////////
	/////////////////////////////////TODO/////////////////////////////////
	//////////////////////////////////////////////////////////////////////
	m.App.Session.Put(r.Context(), "success", "Sub Category updated success")
	http.Redirect(w, r, "/private/admin/dashboard", http.StatusSeeOther)
}

// AdminAddNewService renders adding new service page
func (m *Repository) AdminAddNewService(w http.ResponseWriter, r *http.Request) {

	/////////////////////////////////////////////////////////////////////
	/////////////////////////////////TODO/////////////////////////////////
	//////////////////////////////////////////////////////////////////////

	render.TemplatesRenderer(w, r, "admin-add-new-service.page.tmpl", &models.TemplateData{
		Form:         forms.New(nil),
		Data:         make(map[string]interface{}),
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
	})
}

// PostAdminAddNewService handles post request for adding or editing service info
func (m *Repository) PostAdminAddNewService(w http.ResponseWriter, r *http.Request) {

	/////////////////////////////////////////////////////////////////////
	/////////////////////////////////TODO/////////////////////////////////
	//////////////////////////////////////////////////////////////////////

	render.TemplatesRenderer(w, r, "admin-add-new-service.page.tmpl", &models.TemplateData{
		CategoryMain: m.App.CategoryMain,
		CategorySub:  m.App.CategorySub,
		Services:     m.App.Services,
	})
}
