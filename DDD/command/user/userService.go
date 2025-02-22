package user

import "DDD/entities"

// domain service
type UserService struct {
	userRepository IUserRepository
}
func NewUserService(r UserRepository) *UserService {
	return &UserService{
		userRepository: &r,
	}
}

func (s *UserService) Create(id int, name string) error {
	user, err := entities.NewUser(id, name)
	if err != nil {
		return err
	}
	if err:=s.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}

type IUserRepository interface {
	Create(id int) (entities.User ,error)
	Save(user entities.User)error
}


