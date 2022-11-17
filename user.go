package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"` //binding:"required"
	Username string `json:"username"`
	Password string `json:"password"`
}
