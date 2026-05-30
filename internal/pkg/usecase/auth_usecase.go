package usecase

import (
	"errors"
	"time"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userRepo  repository.UserRepository
	storeRepo repository.StoreRepository
}

func NewAuthUseCase(userRepo repository.UserRepository, storeRepo repository.StoreRepository) AuthUseCase {
	return &authUseCase{
		userRepo:  userRepo,
		storeRepo: storeRepo,
	}
}

func (u *authUseCase) Register(req model.RegisterRequest) (*model.AuthResponse, error) {
	// Check if email already exists
	existingUser, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Check if phone already exists
	existingPhone, err := u.userRepo.GetByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if existingPhone != nil {
		return nil, errors.New("phone already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	now := time.Now().Unix()
	user := &entity.User{
		Email:     req.Email,
		Phone:     req.Phone,
		Name:      req.Name,
		Password:  string(hashedPassword),
		IsAdmin:   false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Auto create store for user
	storeName := req.Name + "'s Store"
	store := &entity.Store{
		UserID:    user.ID,
		Name:      storeName,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.storeRepo.Create(store); err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token: token,
		User: model.UserResponse{
			ID:      user.ID,
			Email:   user.Email,
			Phone:   user.Phone,
			Name:    user.Name,
			IsAdmin: user.IsAdmin,
		},
	}, nil
}

func (u *authUseCase) Login(req model.LoginRequest) (*model.AuthResponse, error) {
	// Get user by email
	user, err := u.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token: token,
		User: model.UserResponse{
			ID:      user.ID,
			Email:   user.Email,
			Phone:   user.Phone,
			Name:    user.Name,
			IsAdmin: user.IsAdmin,
		},
	}, nil
}
