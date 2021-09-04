package after

type User struct {
	id    int
	name  string
	phone int
}

func (u *User) GetSender() *User {
	return &User{
		id:    1,
		name:  "Malik",
		phone: 6281,
	}
}

func (u *User) GetReceiver() *User {
	return &User{
		id:    2,
		name:  "Fajrian",
		phone: 6287,
	}
}
