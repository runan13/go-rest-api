package types

type PostLocalesRequest struct {
	Name string `json:"name"`
}

type PostRestaurantRequest struct {
	Name     string `json:"Name" validate:"required"`
	Location string `json:"Location" validate:"required"`
	Rating   int    `json:"Rating" validate:"required,number"`
}
