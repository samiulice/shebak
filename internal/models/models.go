package models

import (
	"time"
)

// User defines the Admin model with various access level
// AccessLevel defines the admin superiority
// AccessLevel := 1 means SuperAdmin
// AccessLevel := 2 means Admin
// AccessLevel := 3 means Employee
// AccessLevel := 4 means User
type Admin struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// ServiceCategoryMain defines Main Category model
type ServiceCategoryMain struct {
	ID          int
	Name        string
	Available   int
	Description string
	Thumbnanil  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//ServiceCategorySub defines Sub Category model
type ServiceCategorySub struct {
	ID          int
	Name        string
	Available   int
	CategoryID  int
	Description string
	Thumbnanil  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//Service defines service model
type Service struct {
	ID          int
	Name        string
	Available   int
	MinimumCharge int
	CategoryID  int
	SubCategoryID  int
	Country string
	Division string
	District string
	City string
	Description string
	Thumbnanil  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// MailData holds mail messages info
type MailData struct {
	From     string
	To       string
	Subject  string
	Content  string
	Template string
}
