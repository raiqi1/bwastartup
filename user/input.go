package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"` // `json:"name"` is used to change the name of the field when it is returned as a JSON response
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"` // `binding:"required,email"` is used to validate the email field
	Password   string `json:"password" binding:"required"`
}
