package usecase

import (
	"errors"
	"time"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/repository"
)

type storeUseCase struct {
	storeRepo repository.StoreRepository
	userRepo  repository.UserRepository
}

func NewStoreUseCase(storeRepo repository.StoreRepository, userRepo repository.UserRepository) StoreUseCase {
	return &storeUseCase{
		storeRepo: storeRepo,
		userRepo:  userRepo,
	}
}

func (u *storeUseCase) CreateStore(userID uint, req model.CreateStoreRequest) (*model.StoreResponse, error) {
	// Check if user exists
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Check if user already has a store
	existingStore, err := u.storeRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if existingStore != nil {
		return nil, errors.New("user already has a store")
	}

	now := time.Now().Unix()
	store := &entity.Store{
		UserID:    userID,
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.storeRepo.Create(store); err != nil {
		return nil, err
	}

	return &model.StoreResponse{
		ID:      store.ID,
		UserID:  store.UserID,
		Name:    store.Name,
		Address: store.Address,
		Phone:   store.Phone,
	}, nil
}

func (u *storeUseCase) GetStoreByUserID(userID uint) (*model.StoreResponse, error) {
	store, err := u.storeRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("store not found")
	}

	return &model.StoreResponse{
		ID:      store.ID,
		UserID:  store.UserID,
		Name:    store.Name,
		Address: store.Address,
		Phone:   store.Phone,
	}, nil
}

func (u *storeUseCase) GetStoreByID(id uint) (*model.StoreResponse, error) {
	store, err := u.storeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("store not found")
	}

	return &model.StoreResponse{
		ID:      store.ID,
		UserID:  store.UserID,
		Name:    store.Name,
		Address: store.Address,
		Phone:   store.Phone,
	}, nil
}

func (u *storeUseCase) GetAllStores(limit, offset int, name string) (*model.GetStoresResponse, error) {
	filter := entity.FilterStore{
		Limit:  limit,
		Offset: offset,
		Name:   name,
	}

	stores, total, err := u.storeRepo.GetAll(filter)
	if err != nil {
		return nil, err
	}

	var responses []model.StoreResponse
	for _, store := range stores {
		responses = append(responses, model.StoreResponse{
			ID:      store.ID,
			UserID:  store.UserID,
			Name:    store.Name,
			Address: store.Address,
			Phone:   store.Phone,
		})
	}

	return &model.GetStoresResponse{
		Data:  responses,
		Total: total,
	}, nil
}

func (u *storeUseCase) UpdateStore(userID, storeID uint, req model.UpdateStoreRequest) (*model.StoreResponse, error) {
	store, err := u.storeRepo.GetByID(storeID)
	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("store not found")
	}

	// Check authorization: only store owner can update
	if store.UserID != userID {
		return nil, errors.New("unauthorized: you can only update your own store")
	}

	if req.Name != "" {
		store.Name = req.Name
	}
	if req.Address != "" {
		store.Address = req.Address
	}
	if req.Phone != "" {
		store.Phone = req.Phone
	}

	store.UpdatedAt = time.Now().Unix()

	if err := u.storeRepo.Update(store); err != nil {
		return nil, err
	}

	return &model.StoreResponse{
		ID:      store.ID,
		UserID:  store.UserID,
		Name:    store.Name,
		Address: store.Address,
		Phone:   store.Phone,
	}, nil
}

func (u *storeUseCase) DeleteStore(userID, storeID uint) error {
	store, err := u.storeRepo.GetByID(storeID)
	if err != nil {
		return err
	}

	if store == nil {
		return errors.New("store not found")
	}

	// Check authorization
	if store.UserID != userID {
		return errors.New("unauthorized: you can only delete your own store")
	}

	return u.storeRepo.Delete(storeID)
}
