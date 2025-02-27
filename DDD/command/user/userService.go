package user

import (
	"DDD/entities"
)

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
	if err := s.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(id string, name string) error {
	user, err := s.userRepository.GetByID(id) // TODO: replace Create with GetByID
	if err != nil {
		return err
	}
	userName, err := entities.NewUserName(name)
	if err != nil {
		return err
	}
	user.Name = userName
	if err := s.userRepository.Save(user); err != nil {
		return err
	}
	return nil
}

type IUserRepository interface {
	Create(id int) (entities.User, error)
	Save(user entities.User) error
	Update(id string, name string) error
	GetByID(id string) (entities.User, error) // Added GetByID method
	Delete(id uint)error
}
