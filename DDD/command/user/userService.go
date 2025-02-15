package user

import "DDD/entities"

// domain service
type UserService struct {
}

type IUserRepository interface {
	Save(user entities.User)error
}


