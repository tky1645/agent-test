package user

import "DDD/entities"

var _ IUserRepository = (*UserRepository)(nil)

type userTable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user entities.User) error {
	return nil
}

func (r *UserRepository) Create(id int) entities.User {
	user := r.get(id)
	return entities.NewUser(user.ID, user.Name)
}

func (r *UserRepository) get(id int) userTable{
	return userTable{
		ID:   id,
		Name: "getJohn",
	}
}