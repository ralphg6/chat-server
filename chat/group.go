package chat

type Group struct {
	Id    string
	Name  string
	Users map[string]User
}
