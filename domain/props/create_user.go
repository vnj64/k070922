package props

import (
	"project/domain/models"
	"project/shared/dto"
)

type CreateUserReq struct {
	Email      string
	FirstName  *string
	SecondName *string
}

type CreateUserResp struct {
	User *models.User `json:"user"`
}

func DtoToProps(dto dto.CreateUserDto) CreateUserReq {
	return CreateUserReq{
		Email:      dto.Email,
		FirstName:  dto.FirstName,
		SecondName: dto.SecondName,
	}
}
