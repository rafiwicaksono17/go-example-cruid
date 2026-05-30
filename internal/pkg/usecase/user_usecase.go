package usecase

import (
	"errors"
	"time"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/repository"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) GetUserByID(id uint) (*model.UserResponse, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return &model.UserResponse{
		ID:      user.ID,
		Email:   user.Email,
		Phone:   user.Phone,
		Name:    user.Name,
		IsAdmin: user.IsAdmin,
	}, nil
}

func (u *userUseCase) GetAllUsers(limit, offset int, name string) (*model.GetUsersResponse, error) {
	filter := entity.FilterUser{
		Limit:  limit,
		Offset: offset,
		Name:   name,
	}

	users, total, err := u.userRepo.GetAll(filter)
	if err != nil {
		return nil, err
	}

	var responses []model.UserResponse
	for _, user := range users {
		responses = append(responses, model.UserResponse{
			ID:      user.ID,
			Email:   user.Email,
			Phone:   user.Phone,
			Name:    user.Name,
			IsAdmin: user.IsAdmin,
		})
	}

	return &model.GetUsersResponse{
		Data:  responses,
		Total: total,
	}, nil
}

func (u *userUseCase) UpdateUser(id uint, req model.UpdateUserRequest) (*model.UserResponse, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Check if email is already taken
	if req.Email != "" && req.Email != user.Email {
		existingUser, err := u.userRepo.GetByEmail(req.Email)
		if err != nil {
			return nil, err
		}
		if existingUser != nil {
			return nil, errors.New("email already registered")
		}
		user.Email = req.Email
	}

	// Check if phone is already taken
	if req.Phone != "" && req.Phone != user.Phone {
		existingPhone, err := u.userRepo.GetByPhone(req.Phone)
		if err != nil {
			return nil, err
		}
		if existingPhone != nil {
			return nil, errors.New("phone already registered")
		}
		user.Phone = req.Phone
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	user.UpdatedAt = time.Now().Unix()

	if err := u.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:      user.ID,
		Email:   user.Email,
		Phone:   user.Phone,
		Name:    user.Name,
		IsAdmin: user.IsAdmin,
	}, nil
}

func (u *userUseCase) DeleteUser(id uint) error {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	return u.userRepo.Delete(id)
}
