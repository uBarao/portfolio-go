package models

type Project struct {
	Id           int
	Name         string
	Summary      string
	Description  string
	Technologies []string
	Image        string
}
