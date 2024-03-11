package web

type CategoryCreateRequest struct {
	Name string `validate:"required,min=5,max=100" json:"name"`
}
