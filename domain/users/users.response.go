package users

import (
	"app/database/db_scopes"
	"app/database/models"
	"time"
)

type ResponseType map[string]any

type UserShort struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserShortTransform(user models.User) UserShort {
	return UserShort{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func UsersShortTransform(users []models.User) []UserShort {
	var data = []UserShort{}

	for _, user := range users {
		data = append(data, UserShortTransform(user))
	}

	return data
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UserTransform(user models.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

func UsersTransform(users []models.User) []User {
	var data = []User{}

	for _, user := range users {
		data = append(data, UserTransform(user))
	}

	return data
}

type UserResponseType struct {
	User User `json:"user"`
}

func UserResponse(user models.User) ResponseType {
	return ResponseType{
		"user": UserTransform(user),
	}
}

type UsersResponseType struct {
	Users []User `json:"users"`
}

func UsersResponse(users []models.User) ResponseType {
	return ResponseType{
		"users": UsersTransform(users),
	}
}

type UsersMetaResponseType struct {
	Users []User                     `json:"users"`
	Meta  db_scopes.PaginateMetadata `json:"meta"`
}

func UsersMetaResponse(users []models.User, meta db_scopes.PaginateMetadata) ResponseType {
	return ResponseType{
		"users": UsersTransform(users),
		"meta":  meta,
	}
}
