package MessageTypes

type Profile struct {
	Name    string
	Hobbies []string
}

type UserToken struct {
	MessageName string
	Token       string
}
