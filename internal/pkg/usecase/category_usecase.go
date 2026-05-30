package usecase

import (
	"errors"
	"time"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/repository"
)

type categoryUseCase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{categoryRepo: categoryRepo}
}

func (u *categoryUseCase) CreateCategory(req model.CreateCategoryRequest) (*model.CategoryResponse, error) {
	// Check if category already exists
	existing, err := u.categoryRepo.GetByName(req.Name)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("category already exists")
	}

	now := time.Now().Unix()
	category := &entity.Category{
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	return &model.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (u *categoryUseCase) GetCategoryByID(id uint) (*model.CategoryResponse, error) {
	category, err := u.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	return &model.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (u *categoryUseCase) GetAllCategories(limit, offset int, name string) (*model.GetCategoriesResponse, error) {
	filter := entity.FilterCategory{
		Limit:  limit,
		Offset: offset,
		Name:   name,
	}

	categories, total, err := u.categoryRepo.GetAll(filter)
	if err != nil {
		return nil, err
	}

	var responses []model.CategoryResponse
	for _, category := range categories {
		responses = append(responses, model.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return &model.GetCategoriesResponse{
		Data:  responses,
		Total: total,
	}, nil
}

func (u *categoryUseCase) UpdateCategory(id uint, req model.UpdateCategoryRequest) (*model.CategoryResponse, error) {
	category, err := u.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	// Check if new name already exists
	if req.Name != category.Name {
		existing, err := u.categoryRepo.GetByName(req.Name)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, errors.New("category name already exists")
		}
	}

	category.Name = req.Name
	category.UpdatedAt = time.Now().Unix()

	if err := u.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return &model.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (u *categoryUseCase) DeleteCategory(id uint) error {
	category, err := u.categoryRepo.GetByID(id)
	if err != nil {
		return err
	}

	if category == nil {
		return errors.New("category not found")
	}

	return u.categoryRepo.Delete(id)
}
