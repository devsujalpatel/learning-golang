package user

type User struct {
	Name string
	Age  string
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() string {
	return u.Age
}