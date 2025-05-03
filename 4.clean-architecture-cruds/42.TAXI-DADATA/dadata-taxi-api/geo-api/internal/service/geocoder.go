// internal/service/geocoder.go
package service

import (
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/entity"
	"GoPackMVP2025/42.TAXI-DADATA/dadata-taxi-api/geo-api/internal/repository"
	"context"

	"github.com/ekomobile/dadata/v2/api/suggest"
)

type GeocoderService interface {
	Geocode(ctx context.Context, address string, userID uint) (*entity.Address, error)
	GetHistory(ctx context.Context, userID uint) ([]entity.Address, error)
	DeleteFromHistory(ctx context.Context, id uint, userID uint) error
}

type geocoderService struct {
	addressRepo repository.AddressRepository
	daDataAPI   *dadata.SuggestApi
}

func NewGeocoderService(repo repository.AddressRepository, apiKey, secretKey string) GeocoderService {
	api := dadata.NewSuggestApi()
	// Настройка API с ключами
	return &geocoderService{
		addressRepo: repo,
		daDataAPI:   api,
	}
}

func (s *geocoderService) Geocode(ctx context.Context, address string, userID uint) (*entity.Address, error) {
	// Вызов DaData API
	params := suggest.RequestParams{Query: address}
	suggestions, err := s.daDataAPI.Address(ctx, &params)
	if err != nil {
		return nil, err
	}

	if len(suggestions) == 0 {
		return nil, ErrAddressNotFound
	}

	// Преобразование результата в нашу модель
	result := suggestions[0]
	geoAddress := &entity.Address{
		UserID:       userID,
		RawAddress:   address,
		Result:       result.Value,
		PostalCode:   result.Data.PostalCode,
		Country:      result.Data.Country,
		Region:       result.Data.Region,
		CityArea:     result.Data.CityArea,
		CityDistrict: result.Data.CityDistrict,
		Street:       result.Data.Street,
		House:        result.Data.House,
		GeoLat:       result.Data.GeoLat,
		GeoLon:       result.Data.GeoLon,
		QCGeo:        result.Data.QCGeo,
	}

	// Сохранение в БД
	if err := s.addressRepo.Create(ctx, geoAddress); err != nil {
		return nil, err
	}

	return geoAddress, nil
}

// Реализация остальных методов...
