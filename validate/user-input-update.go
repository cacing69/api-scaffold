package validate

type UserInputUpdate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
