package userMapper

import (
	shared "go-rest-template/api/shared/domain"
	user "go-rest-template/api/user/domain"
	userValueObjects "go-rest-template/api/user/domain/valueObjects"
	"go-rest-template/internal/app/utils"
)

func ToDomain(rawUser user.RawUser) user.User {
	createdAt, _ := utils.StringToTime(rawUser.CreatedAt)
	updatedAt, _ := utils.StringToTime(rawUser.UpdatedAt)

	return user.User{
		ID:        shared.ID{Value: rawUser.ID},
		UUID:      shared.UUID{Value: rawUser.UUID},
		CreatedAt: shared.Date{Value: createdAt},
		UpdatedAt: shared.Date{Value: updatedAt},
		Name:      userValueObjects.Name{Value: rawUser.Name},
		Email:     userValueObjects.Email{Value: rawUser.Email},
		Password:  userValueObjects.Password{Value: rawUser.Password},
	}
}
