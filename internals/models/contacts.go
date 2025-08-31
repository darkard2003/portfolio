package models

type Handle struct {
	Handle string `json:"handle"`
	Url    string `json:"url"`
}

type Contacts struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	LinkedIn Handle `json:"linkedin"`
	Github   Handle `json:"github"`
}
