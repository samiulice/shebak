package models

import (
	"github.com/samiulru/shebak/internal/forms"
)

// TemplateData holds data sent from handler to template
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CategoryMain    []ServiceCategoryMain
	CategorySub     []ServiceCategorySub
	Services        []Service
	User            string
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
