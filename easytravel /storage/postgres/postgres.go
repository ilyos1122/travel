package postgres

import (
	"database/sql"
	"fmt"
	"app/storage"
	"app/config"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sql.DB
	country storage.CountryRepoI
	city    storage.CityRepoI
	aeroport storage.AeroportRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) Country() storage.CountryRepoI {

	if s.country == nil {
		s.country = NewCountryRepo(s.db)
	}

	return s.country
}

func (s *Store) City() storage.CityRepoI {

	if s.city == nil {
		s.city = NewCityRepo(s.db)
	}

	return s.city
}

func (s *Store) Aeroport() storage.AeroportRepoI {

	if s.aeroport == nil {
		s.aeroport = NewAeroportRepo(s.db)
	}

	return s.aeroport
}
