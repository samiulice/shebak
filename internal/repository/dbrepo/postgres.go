package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/samiulru/shebak/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// GetAllServiceCategoryMain returns a slice of all Service Category Main from the database
func (m *postgresDBRepo) GetAllServiceCategoryMain() ([]models.ServiceCategoryMain, error) {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var allMainCategory []models.ServiceCategoryMain
	query := `select id, name, available ,description, thumbnail, created_at, updated_at from service_category_main order by id`

	rows, err := m.DB.QueryContext(cntx, query)
	if err != nil {
		return allMainCategory, err
	}
	defer rows.Close()

	for rows.Next() {
		var mainCategory models.ServiceCategoryMain
		err = rows.Scan(
			&mainCategory.ID,
			&mainCategory.Name,
			&mainCategory.Available,
			&mainCategory.Description,
			&mainCategory.Thumbnanil,
			&mainCategory.CreatedAt,
			&mainCategory.UpdatedAt,
		)
		if err != nil {
			return allMainCategory, err
		}
		allMainCategory = append(allMainCategory, mainCategory)
	}
	if err = rows.Err(); err != nil {
		return allMainCategory, err
	}
	return allMainCategory, nil
}

// GetServiceCategoryMainByID searches and returns Service Category Main info by ID from the database
func (m *postgresDBRepo) GetServiceCategoryMainByID(id int) (models.ServiceCategoryMain, error) {

	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var categoryMain models.ServiceCategoryMain
	query := `select id, name, available, description, thumbnail, created_at, updated_at
		from service_category_main where id = $1`

	row := m.DB.QueryRowContext(cntx, query, id)
	err := row.Scan(
		&categoryMain.ID,
		&categoryMain.Name,
		&categoryMain.Available,
		&categoryMain.Description,
		&categoryMain.Thumbnanil,
		&categoryMain.CreatedAt,
		&categoryMain.UpdatedAt,
	)

	if err != nil {
		return categoryMain, err
	}

	return categoryMain, nil
}


// UpdateServiceCategoryMain existing Service Category Main info to the database
func (m *postgresDBRepo) UpdateServiceCategoryMain(categoryMain models.ServiceCategoryMain) error {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update service_category_main
			set name = $1 , available = $2 , description = $3 , thumbanil = $4 , updated_at = $5
			where id = $6`

	_, err := m.DB.ExecContext(cntx, query,
		categoryMain.Name,
		categoryMain.Available,
		categoryMain.Description,
		categoryMain.Thumbnanil,
		time.Now(),
		categoryMain.ID,
	)

	return err
}

// InsertServiceCategoryMain inserts CategoryMain info to the database
func (m *postgresDBRepo) InsertServiceCategoryMain(service models.ServiceCategoryMain) (int, error) {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	query := `insert into service_category_main (name, description, thumbnail, created_at, updated_at)
			values ($1, $2, $3, $4, $5) returning id`

	err := m.DB.QueryRowContext(cntx, query,
		&service.Name,
		&service.Description,
		&service.Thumbnanil,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

/////////////////Service Category Sub/////////////////////////////////
// InsertServiceCategorySub inserts CategorySub info to the database
func (m *postgresDBRepo) InsertServiceCategorySub(service models.ServiceCategorySub) (int, error) {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	query := `insert into service_category_sub (name, category_id, description, thumbnail, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6) returning id`

	err := m.DB.QueryRowContext(cntx, query,
		&service.Name,
		&service.CategoryID,
		&service.Description,
		&service.Thumbnanil,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// GetAllServiceCategorySub returns a slice of all Sub Categories from the database
func (m *postgresDBRepo) GetAllServiceCategorySub() ([]models.ServiceCategorySub, error) {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var allSubCategory []models.ServiceCategorySub
	query := `select id, name, available, category_id ,description, thumbnail, created_at, updated_at from service_category_sub order by name`

	rows, err := m.DB.QueryContext(cntx, query)
	if err != nil {
		return allSubCategory, err
	}
	defer rows.Close()

	for rows.Next() {
		var subCategory models.ServiceCategorySub
		err = rows.Scan(
			&subCategory.ID,
			&subCategory.Name,
			&subCategory.Available,
			&subCategory.CategoryID,
			&subCategory.Description,
			&subCategory.Thumbnanil,
			&subCategory.CreatedAt,
			&subCategory.UpdatedAt,
		)
		if err != nil {
			return allSubCategory, err
		}
		allSubCategory = append(allSubCategory, subCategory)
	}
	if err = rows.Err(); err != nil {
		return allSubCategory, err
	}
	return allSubCategory, nil
}

// GetServiceCategorySubByID searches and returns Service Category Sub info by ID from the database
func (m *postgresDBRepo) GetServiceCategorySubByID(id int) (models.ServiceCategorySub, error) {

	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var categorySub models.ServiceCategorySub
	query := `select id, name, available, category_id, description, thumbnail, created_at, updated_at
		from service_category_sub where id = $1`

	row := m.DB.QueryRowContext(cntx, query, id)
	err := row.Scan(
		&categorySub.ID,
		&categorySub.Name,
		&categorySub.Available,
		&categorySub.CategoryID,
		&categorySub.Description,
		&categorySub.Thumbnanil,
		&categorySub.CreatedAt,
		&categorySub.UpdatedAt,
	)

	if err != nil {
		return categorySub, err
	}

	return categorySub, nil
}

//GetSubListByMainID returns a slice of Service Category Sub based on foreign key category_id
func (m *postgresDBRepo) GetSubListByMainID(category_id int) ([]models.ServiceCategorySub, error){
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var SubList []models.ServiceCategorySub
	query := `select id, name, available, category_id, description, thumbnail, created_at, updated_at 
				from service_category_sub where category_id = $1 order by id`

	rows, err := m.DB.QueryContext(cntx, query, category_id)
	if err != nil {
		return SubList, err
	}
	defer rows.Close()

	for rows.Next() {
		var sub models.ServiceCategorySub
		err = rows.Scan(
			&sub.ID,
			&sub.Name,
			&sub.Available,
			&sub.CategoryID,
			&sub.Description,
			&sub.Thumbnanil,
			&sub.CreatedAt,
			&sub.UpdatedAt,
		)
		if err != nil {
			return SubList, err
		}
		SubList = append(SubList, sub)
	}
	if err = rows.Err(); err != nil {
		return SubList, err
	}
	return SubList, nil
}

/////////////////////Service///////////////////////////////////////
// GetAllServices returns a slice of all Services from the database
func (m *postgresDBRepo) GetAllServices() ([]models.Service, error) {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var allServices []models.Service
	query := `select id, name, available ,description, thumbnail, created_at, updated_at from service_category_main order by name`

	rows, err := m.DB.QueryContext(cntx, query)
	if err != nil {
		return allServices, err
	}
	defer rows.Close()

	for rows.Next() {
		var service models.Service
		err = rows.Scan(
			&service.ID,
			&service.Name,
			&service.Available,
			&service.Description,
			&service.Thumbnanil,
			&service.CreatedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			return allServices, err
		}
		allServices = append(allServices, service)
	}
	if err = rows.Err(); err != nil {
		return allServices, err
	}
	return allServices, nil
}
////////////////////////////////////////edit this///////////////////////
/////////////////////////////////////////////////////////////////////////
//GetServiceListByMainID_SubID returns a slice of Services based on foreign keys category_id and sub_category_id
func (m *postgresDBRepo) GetServiceListByMainID_SubID(category_id, sub_category_id int) ([]models.Service, error){
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var ServiceList []models.Service
	query := `select id, name, available, minimum_charge, category_id, sub_category_id, description, thumbnail, country, division, district, city, created_at, updated_at from service where category_id = $1 and sub_category_id = $2`

	rows, err := m.DB.QueryContext(cntx, query, category_id, sub_category_id)
	if err != nil {
		return ServiceList, err
	}
	defer rows.Close()

	for rows.Next() {
		var service models.Service
		err = rows.Scan(
			&service.ID,
			&service.Name,
			&service.Available,
			&service.MinimumCharge,
			&service.CategoryID,
			&service.SubCategoryID,
			&service.Description,
			&service.Thumbnanil,
			&service.Country,
			&service.Division,
			&service.District,
			&service.City,
			&service.CreatedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			return ServiceList, err
		}
		ServiceList = append(ServiceList, service)
	}
	if err = rows.Err(); err != nil {
		return ServiceList, err
	}
	return ServiceList, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////Deprecated:: Old functions needs to be updated////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// AllUsers retrive all users from database
func (m *postgresDBRepo) AllAdmins() bool {
	return true
}

// GetUserByID searches user by ID
func (m *postgresDBRepo) GetAdminByID(id int) (models.Admin, error) {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.Admin
	query := `select id, first_name, last_name, email, password, access_level, created_at, updated_at
		from users where id = $1`

	row := m.DB.QueryRowContext(cntx, query, id)
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdateAt,
	)

	if err != nil {
		return u, err
	}
	return u, nil
}

// UpdateUser updates a user in the database
func (m *postgresDBRepo) UpdateAdmin(u models.Admin) error {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `update users
			set first_name = $1 , last_name = $2 , email = $3 , password = $4 , access_level = $5, updated_at = $6
			where id = $7`

	_, err := m.DB.ExecContext(cntx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.AccessLevel,
		time.Now(),
		u.ID,
	)

	return err
}

// Authenticate authenticates a user
func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, int, error) {
	cntx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id, accessLevel int
	var hashedPassword string

	row := m.DB.QueryRowContext(cntx, "select id, password, access_level from admin where email = $1", email)

	err := row.Scan(&id, &hashedPassword, &accessLevel)
	if err != nil {
		return id, "", accessLevel, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return 0, "", 0, errors.New("incorrect password")
	} else if err != nil {
		return 0, "", 0, err
	}

	return id, hashedPassword, accessLevel, nil
}
