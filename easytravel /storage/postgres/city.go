package postgres

import (
	"database/sql"
	"fmt"

	"app/models"
	"github.com/google/uuid"
)

type cityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) *cityRepo {
	return &cityRepo{
		db: db,
	}
}

func (r *cityRepo) Create(req *models.CreateCity) (*models.City, error) {

	var (
		cityId = uuid.New().String()
		query      = `
			INSERT INTO "city"(
				"id",
				"title",
				"country_id",
				"city_code",
				"latitude",
				"longitude",
				"offset",
				"country_name",
				"updated_at"
			) VALUES ($1, $2, $3, $4,$5, $6, $7, $8, NOW())`
	)

	_, err := r.db.Exec(
		query,
		cityId,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.CountryName,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.CityPrimaryKey{Id: cityId})
}

func (r *cityRepo) GetByID(req *models.CityPrimaryKey) (*models.City, error) {

	var (
		city models.City
		query    = `
			SELECT
					"id",
					"title",
					"country_id",
					"city_code",
					"latitude",
					"longitude",
					"offset",
					"country_name",
					"created_at",
					"updated_at"
			FROM "city"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&city.Id,
		&city.Title,
		&city.CountryId,
		&city.CityCode,
		&city.Latitude,
		&city.Longitude,
		&city.Offset,
		&city.CountryName,
		&city.CreatedAt,
		&city.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &city, nil
}

func (r *cityRepo) GetList(req *models.GetListCityRequest) (*models.GetListCityResponse, error) {
	var (
		resp   models.GetListCityResponse
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
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"country_name",
			"created_at",
			"updated_at"
		FROM "city"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			city models.City
		)

		err = rows.Scan(
			&resp.Count,
			&city.Id,
			&city.Title,
			&city.CountryId,
			&city.CityCode,
			&city.Latitude,
			&city.Longitude,
			&city.Offset,
			&city.CountryName,
			&city.CreatedAt,
			&city.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Cities = append(resp.Cities, &city)
	}

	return &resp, nil
}

func (r *cityRepo) Update(req *models.UpdateCity) (int64, error) {

	query := `
		UPDATE city
			SET
				"title" =$2
				"country_id" =$3
				"city_code" =$4
				"latitude" =$5
				"longitude" =$6
				"offset" =$7
				"country_name" =$8
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.CountryName,
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

func (r *cityRepo) Delete(req *models.CityPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM city WHERE id = $1", req.Id)
	return err
}
