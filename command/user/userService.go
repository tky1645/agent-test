package user

import (
	"DDD/entities"
	"strconv"
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
	user, err := s.userRepository.GetByID(id)
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

func (s *UserService) Delete(id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return s.userRepository.Delete(uint(idUint))
}

func (s *UserService) GetByID(id string) (entities.User, error) {
	return s.userRepository.GetByID(id)
}

type IUserRepository interface {
	Create(id int) (entities.User, error)
	Save(user entities.User) error
	Update(id string, name string) error
	GetByID(id string) (entities.User, error)
	Delete(id uint) error
}
