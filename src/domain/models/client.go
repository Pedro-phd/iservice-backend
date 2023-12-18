package models

type Client struct {
	BasicModel
	Name     string
	Email    string
	Document string
	Phone    string
	Password string
	Image    string
	State    string
	City     string
	Street   string
	Number   string
	Active   bool
}
