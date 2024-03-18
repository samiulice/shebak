package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/samiulru/shebak/internal/models"
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	CategoryMain []models.ServiceCategoryMain
	CategorySub []models.ServiceCategorySub
	Services []models.Service
	InProduction  bool
	Session       *scs.SessionManager
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	MailChan      chan models.MailData
}
