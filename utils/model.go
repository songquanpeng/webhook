package utils

import "gorm.io/gorm"

type Webhook struct {
	gorm.Model
	Id              uint
	Name            string
	Description     string
	ShellScriptPath string
	Url             string
	Secret          string
}
