package controller

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUserById(id string) User {
	return User{
		ID:   id,
		Name: "name",
		Age:  18,
	}
}

func DeleteUserById(id string) User {
	return User{
		ID:   id,
		Name: "name",
		Age:  18,
	}
}

func UpdateUser(id string, user User) User {
	user.ID = id
	return user
}

func NewUser(user User) User {
	user.ID = "id"
	return user
}
