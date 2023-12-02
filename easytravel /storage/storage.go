package storage

import "app/models"

type StorageI interface {
	Country() CountryRepoI
	City() CityRepoI
	Aeroport() AeroportRepoI
}

type CountryRepoI interface {
	Create(req *models.CreateCountry) (*models.Country, error)
	GetByID(req *models.CountryPrimaryKey) (*models.Country, error)
	GetList(req *models.GetListCountryRequest) (*models.GetListCounrtyResponse, error)
	Update(req *models.UpdateCountry) (int64, error)
	Delete(req *models.CountryPrimaryKey) error
}

type CityRepoI interface {
	Create(req *models.CreateCity) (*models.City, error)
	GetByID(req *models.CityPrimaryKey) (*models.City, error)
	GetList(req *models.GetListCityRequest) (*models.GetListCityResponse, error)
	Update(req *models.UpdateCity) (int64, error)
	Delete(req *models.CityPrimaryKey) error
}

type AeroportRepoI interface {
	Create(req *models.CreateAeroport) (*models.Aeroport, error)
	GetByID(req *models.AeroportPrimaryKey) (*models.Aeroport, error)
	GetList(req *models.GetListAeroportRequest) (*models.GetListAeroportResponse, error)
	Update(req *models.UpdateAeroport) (int64, error)
	Delete(req *models.AeroportPrimaryKey) error
}
