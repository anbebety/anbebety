package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name                string
	Telephone           string
	Password            string
	PersonalInformation string
	Identity            int
}
type Group struct {
	gorm.Model
	Name     string
	Title    string
	Aim      string
	Time     string
	Location string
	Require  string
	Number   int
	State    int
}
type Apply struct {
	gorm.Model
	Name       string
	GroupTitle string
	Reason     string
	Advantage  string
	State      int
}
type Message struct {
	gorm.Model
	Sender   string
	Content  string
	Receiver string
	State    int
}
type Session struct {
	gorm.Model
	Name  string
	Value string
}
