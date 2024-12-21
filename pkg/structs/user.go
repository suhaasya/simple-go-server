package structs

type AddUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

type EditUser struct {
	Name     string `json:"name" validate:"omitempty"`
	Email    string `json:"email" validate:"omitempty,email"`
	Role     string `json:"role" validate:"omitempty,oneof=admin user"`
}

type QueryParams struct {
	Page    string `json:"page" validate:"omitempty,numeric,min=1"`
	Size    string `json:"size" validate:"omitempty,numeric,min=1"`
	Order   string `json:"order" validate:"omitempty,oneof=asc desc"`
	OrderBy string `json:"orderBy" validate:"omitempty,oneof=id name"`
}
