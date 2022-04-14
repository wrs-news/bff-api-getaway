// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type NewUser struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StatusResp struct {
	Status string `json:"status"`
}

type Tokens struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

type UpdateUser struct {
	UUID  string `json:"uuid"`
	Login string `json:"login"`
	Email string `json:"email"`
	Role  int    `json:"role"`
}

type User struct {
	UUID      string `json:"uuid"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserSelection struct {
	Limit    int     `json:"limit"`
	Offset   int     `json:"offset"`
	Data     []*User `json:"data"`
	Total    int     `json:"total"`
	LastPage int     `json:"lastPage"`
}
