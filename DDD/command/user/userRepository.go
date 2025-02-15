package user

import "DDD/entities"

var _ IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user entities.User) error {
	return nil
}

func (r *UserRepository) Create(id int) entities.User {
	 return entities.NewUser(id, "createJohn")
}