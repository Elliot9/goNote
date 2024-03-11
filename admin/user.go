package admin

type User struct {
	Account  string `form:"account"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Sex      int    `form:"sex"`
}

func (u User) TableName() string {
	return "users"
}
