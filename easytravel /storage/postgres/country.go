package postgres

import (
	"database/sql"
	"fmt"

	"app/models"
	"github.com/google/uuid"
)

type countryRepo struct {
	db *sql.DB
}

func NewCountryRepo(db *sql.DB) *countryRepo {
	return &countryRepo{
		db: db,
	}
}

func (r *countryRepo) Create(req *models.CreateCountry) (*models.Country, error) {

	var (
		countryId = uuid.New().String()
		query      = `
			INSERT INTO "country"(
				"id",
				"title",
				"code",
				"continent",
				"updated_at"
			) VALUES ($1, $2, $3, $4, NOW())`
	)

	_, err := r.db.Exec(
		query,
		countryId,
		req.Title,
		req.Code,
		req.Continent,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.CountryPrimaryKey{Id: countryId})
}

func (r *countryRepo) GetByID(req *models.CountryPrimaryKey) (*models.Country, error) {

	var (
		country models.Country
		query    = `
			SELECT
				"id",
				"title",
				"code",
				"continent",
				"created_at",
				"updated_at"	
			FROM "country"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&country.Id,
		&country.Title,
		&country.Code,
		&country.Continent,
		&country.CreatedAt,
		&country.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (r *countryRepo) GetList(req *models.GetListCountryRequest) (*models.GetListCounrtyResponse, error) {
	var (
		resp   models.GetListCounrtyResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.Search) > 0 {
		where += " AND title ILIKE" + " '%" + req.Search + "%'"
	}

	if len(req.Query) > 0 {
		where += req.Query
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"title",
			"code",
			"continent",
			"created_at",
			"updated_at"
		FROM "country"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			country models.Country
		)

		err = rows.Scan(
			&resp.Count,
			&country.Id,
			&country.Title,
			&country.Code,
			&country.Continent,
			&country.CreatedAt,
			&country.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Countries = append(resp.Countries, &country)
	}

	return &resp, nil
}

func (r *countryRepo) Update(req *models.UpdateCountry) (int64, error) {

	query := `
		UPDATE country
			SET
				title = $2,
				code = $3,
				continent = $4
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.Title,
		req.Code,
		req.Continent,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *countryRepo) Delete(req *models.CountryPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM country WHERE id = $1", req.Id)
	return err
}
