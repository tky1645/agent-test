package user

import "DDD/entities"

// domain service
type UserService struct {
}

type IUserRepository interface {
	Create(id int) entities.User
	Save(user entities.User)error
}


