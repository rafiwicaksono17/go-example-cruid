package usecase

import (
	"errors"
	"time"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/repository"
)

type addressUseCase struct {
	addressRepo repository.AddressRepository
	userRepo    repository.UserRepository
}

func NewAddressUseCase(addressRepo repository.AddressRepository, userRepo repository.UserRepository) AddressUseCase {
	return &addressUseCase{
		addressRepo: addressRepo,
		userRepo:    userRepo,
	}
}

func (u *addressUseCase) CreateAddress(userID uint, req model.CreateAddressRequest) (*model.AddressResponse, error) {
	// Check if user exists
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	now := time.Now().Unix()
	address := &entity.Address{
		UserID:        userID,
		JudulAlamat:   req.JudulAlamat,
		PenerimaNama:  req.PenerimaNama,
		PenerimaPhone: req.PenerimaPhone,
		Provinsi:      req.Provinsi,
		ProvinsiID:    req.ProvinsiID,
		Kabupaten:     req.Kabupaten,
		KabupatenID:   req.KabupatenID,
		Kecamatan:     req.Kecamatan,
		KecamatanID:   req.KecamatanID,
		Kelurahan:     req.Kelurahan,
		KelurahanID:   req.KelurahanID,
		DetailAlamat:  req.DetailAlamat,
		IsDefault:     req.IsDefault,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := u.addressRepo.Create(address); err != nil {
		return nil, err
	}

	return &model.AddressResponse{
		ID:            address.ID,
		UserID:        address.UserID,
		JudulAlamat:   address.JudulAlamat,
		PenerimaNama:  address.PenerimaNama,
		PenerimaPhone: address.PenerimaPhone,
		Provinsi:      address.Provinsi,
		ProvinsiID:    address.ProvinsiID,
		Kabupaten:     address.Kabupaten,
		KabupatenID:   address.KabupatenID,
		Kecamatan:     address.Kecamatan,
		KecamatanID:   address.KecamatanID,
		Kelurahan:     address.Kelurahan,
		KelurahanID:   address.KelurahanID,
		DetailAlamat:  address.DetailAlamat,
		IsDefault:     address.IsDefault,
	}, nil
}

func (u *addressUseCase) GetAddressByID(userID, addressID uint) (*model.AddressResponse, error) {
	address, err := u.addressRepo.GetByUserIDAndID(userID, addressID)
	if err != nil {
		return nil, err
	}

	if address == nil {
		return nil, errors.New("address not found")
	}

	return &model.AddressResponse{
		ID:            address.ID,
		UserID:        address.UserID,
		JudulAlamat:   address.JudulAlamat,
		PenerimaNama:  address.PenerimaNama,
		PenerimaPhone: address.PenerimaPhone,
		Provinsi:      address.Provinsi,
		ProvinsiID:    address.ProvinsiID,
		Kabupaten:     address.Kabupaten,
		KabupatenID:   address.KabupatenID,
		Kecamatan:     address.Kecamatan,
		KecamatanID:   address.KecamatanID,
		Kelurahan:     address.Kelurahan,
		KelurahanID:   address.KelurahanID,
		DetailAlamat:  address.DetailAlamat,
		IsDefault:     address.IsDefault,
	}, nil
}

func (u *addressUseCase) GetAllAddressesByUser(userID uint, limit, offset int) (*model.GetAddressesResponse, error) {
	filter := entity.FilterAddress{
		Limit:  limit,
		Offset: offset,
	}

	addresses, total, err := u.addressRepo.GetByUserID(userID, filter)
	if err != nil {
		return nil, err
	}

	var responses []model.AddressResponse
	for _, address := range addresses {
		responses = append(responses, model.AddressResponse{
			ID:            address.ID,
			UserID:        address.UserID,
			JudulAlamat:   address.JudulAlamat,
			PenerimaNama:  address.PenerimaNama,
			PenerimaPhone: address.PenerimaPhone,
			Provinsi:      address.Provinsi,
			ProvinsiID:    address.ProvinsiID,
			Kabupaten:     address.Kabupaten,
			KabupatenID:   address.KabupatenID,
			Kecamatan:     address.Kecamatan,
			KecamatanID:   address.KecamatanID,
			Kelurahan:     address.Kelurahan,
			KelurahanID:   address.KelurahanID,
			DetailAlamat:  address.DetailAlamat,
			IsDefault:     address.IsDefault,
		})
	}

	return &model.GetAddressesResponse{
		Data:  responses,
		Total: total,
	}, nil
}

func (u *addressUseCase) UpdateAddress(userID, addressID uint, req model.UpdateAddressRequest) (*model.AddressResponse, error) {
	address, err := u.addressRepo.GetByUserIDAndID(userID, addressID)
	if err != nil {
		return nil, err
	}

	if address == nil {
		return nil, errors.New("address not found")
	}

	if req.JudulAlamat != "" {
		address.JudulAlamat = req.JudulAlamat
	}
	if req.PenerimaNama != "" {
		address.PenerimaNama = req.PenerimaNama
	}
	if req.PenerimaPhone != "" {
		address.PenerimaPhone = req.PenerimaPhone
	}
	if req.Provinsi != "" {
		address.Provinsi = req.Provinsi
	}
	if req.ProvinsiID != "" {
		address.ProvinsiID = req.ProvinsiID
	}
	if req.Kabupaten != "" {
		address.Kabupaten = req.Kabupaten
	}
	if req.KabupatenID != "" {
		address.KabupatenID = req.KabupatenID
	}
	if req.Kecamatan != "" {
		address.Kecamatan = req.Kecamatan
	}
	if req.KecamatanID != "" {
		address.KecamatanID = req.KecamatanID
	}
	if req.Kelurahan != "" {
		address.Kelurahan = req.Kelurahan
	}
	if req.KelurahanID != "" {
		address.KelurahanID = req.KelurahanID
	}
	if req.DetailAlamat != "" {
		address.DetailAlamat = req.DetailAlamat
	}

	address.IsDefault = req.IsDefault
	address.UpdatedAt = time.Now().Unix()

	if err := u.addressRepo.Update(address); err != nil {
		return nil, err
	}

	return &model.AddressResponse{
		ID:            address.ID,
		UserID:        address.UserID,
		JudulAlamat:   address.JudulAlamat,
		PenerimaNama:  address.PenerimaNama,
		PenerimaPhone: address.PenerimaPhone,
		Provinsi:      address.Provinsi,
		ProvinsiID:    address.ProvinsiID,
		Kabupaten:     address.Kabupaten,
		KabupatenID:   address.KabupatenID,
		Kecamatan:     address.Kecamatan,
		KecamatanID:   address.KecamatanID,
		Kelurahan:     address.Kelurahan,
		KelurahanID:   address.KelurahanID,
		DetailAlamat:  address.DetailAlamat,
		IsDefault:     address.IsDefault,
	}, nil
}

func (u *addressUseCase) DeleteAddress(userID, addressID uint) error {
	address, err := u.addressRepo.GetByUserIDAndID(userID, addressID)
	if err != nil {
		return err
	}

	if address == nil {
		return errors.New("address not found")
	}

	return u.addressRepo.Delete(addressID)
}
