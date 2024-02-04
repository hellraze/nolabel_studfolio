package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type User struct {
	userID       uuid.UUID
	email        string
	password     string
	name         string
	about        string
	education    string
	experience   string
	urls         []string
	imageID      int
	roleID       int
	projectsList []int
}

func (u *User) ID() uuid.UUID       { return u.userID }
func (u *User) Email() string       { return u.email }
func (u *User) Password() string    { return u.password }
func (u *User) Name() string        { return u.name }
func (u *User) About() string       { return u.about }
func (u *User) Education() string   { return u.education }
func (u *User) Experience() string  { return u.experience }
func (u *User) Urls() []string      { return u.urls }
func (u *User) ImageID() int        { return u.imageID }
func (u *User) RoleID() int         { return u.roleID }
func (u *User) ProjectsList() []int { return u.projectsList }

func (u *User) SetName(name string)             { u.name = name }
func (u *User) SetAbout(about string)           { u.about = about }
func (u *User) SetEducation(education string)   { u.education = education }
func (u *User) SetExperience(experience string) { u.experience = experience }
func (u *User) SetImageID(imageID int)          { u.imageID = imageID }
func (u *User) SetRoleID(roleID int)            { u.roleID = roleID }
func (u *User) AddToUrls(url string)            { u.urls = append(u.urls, url) }
func (u *User) AddToProjectsList(projectID int) { u.projectsList = append(u.projectsList, projectID) }

func NewUser(userID uuid.UUID, email string, password string, name string, about string, education string, experience string, urls []string, roleID int, projectsList []int) *User {
	return &User{
		userID:       userID,
		email:        email,
		password:     password,
		name:         name,
		about:        about,
		education:    education,
		experience:   experience,
		urls:         urls,
		roleID:       roleID,
		projectsList: projectsList,
	}
}

type UserRepository interface {
	Save(context.Context, *User) error
	FindByEmail(context.Context, string) (*User, error)
}
