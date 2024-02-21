package user

type Email struct {
	Value string
}

func validateEmail(value string) (Email, error) {
	return Email{Value: value}, nil
}
