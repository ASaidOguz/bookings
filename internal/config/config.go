package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// Appconfig holds the application config
type AppConfig struct {
	UseCache      bool // for development on and off -- very nice element
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
