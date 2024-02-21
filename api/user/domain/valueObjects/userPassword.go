package user

type Password struct {
	Value string
}

func validatePassword(value string) (Password, error) {
	return Password{Value: value}, nil
}
