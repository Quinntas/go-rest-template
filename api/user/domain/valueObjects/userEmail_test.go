package userValueObjects

import (
	"testing"
)

func TestEmptyEmail(t *testing.T) {
	t.Parallel()
	email, err := ValidateEmail("")
	if email != nil {
		t.Errorf("Expected nil, got %v", email)
	}
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestInvalidEmail(t *testing.T) {
	t.Parallel()
	email, err := ValidateEmail("invalidemail")
	if email != nil {
		t.Errorf("Expected nil, got %v", email)
	}
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestValidEmail(t *testing.T) {
	t.Parallel()
	email, err := ValidateEmail("validemail@gmail.com")
	if email == nil {
		t.Errorf("Expected email, got nil")
	}
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
