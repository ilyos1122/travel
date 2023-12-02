package postgres

import (
	"database/sql"
	"fmt"

	"app/models"
	"github.com/google/uuid"
)

type aeroportRepo struct {
	db *sql.DB
}

func NewAeroportRepo(db *sql.DB) *aeroportRepo {
	return &aeroportRepo{
		db: db,
	}
}

func (r *aeroportRepo) Create(req *models.CreateAeroport) (*models.Aeroport, error) {

	var (
		aeroportId = uuid.New().String()
		query      = `
			INSERT INTO "aeroport"(
				"id",
				"title",
				"country_id",
				"city_id",
				"latitude",
				"longitude",
				"radius",
				"image",
				"address",
				"country",
				"city",
				"search_text",
				"code",
				"product_count",
				"gmt",
				"updated_at"
			) VALUES ($1, $2, $3, $4,$5, $6, $7, $8, $9,$10, $11, $12,$13,$14,$15, NOW())`
	)

	_, err := r.db.Exec(
		query,
		aeroportId,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Latitude,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Address,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.AeroportPrimaryKey{Id: aeroportId})
}

func (r *aeroportRepo) GetByID(req *models.AeroportPrimaryKey) (*models.Aeroport, error) {

	var (
		aeroport models.Aeroport
		query    = `
			SELECT
					"id",
					"title",
					"country_id",
					"city_id",
					"latitude",
					"longitude",
					"radius",
					"image",
					"address",
					"country",
					"city",
					"search_text",
					"code",
					"product_count",
					"gmt",
					"created_at",
					"updated_at"
			FROM "aeroport"
			WHERE "id" = $1
		`
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&aeroport.Id,
		&aeroport.Title,
		&aeroport.CountryId,
		&aeroport.CityId,
		&aeroport.Latitude,
		&aeroport.Longitude,
		&aeroport.Radius,
		&aeroport.Image,
		&aeroport.Address,
		&aeroport.Country,
		&aeroport.City,
		&aeroport.SearchText,
		&aeroport.Code,
		&aeroport.ProductCount,
		&aeroport.Gmt,
		&aeroport.CreatedAt,
		&aeroport.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &aeroport, nil
}

func (r *aeroportRepo) GetList(req *models.GetListAeroportRequest) (*models.GetListAeroportResponse, error) {
	var (
		resp   models.GetListAeroportResponse
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
					"city_id",
					"latitude",
					"longitude",
					"radius",
					"image",
					"address",
					"country",
					"city",
					"search_text",
					"code",
					"product_count",
					"gmt",
					"created_at",
					"updated_at"
		FROM "aeroport"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			aeroport models.Aeroport
		)

		err = rows.Scan(
			&resp.Count,
			&aeroport.Id,
			&aeroport.Title,
			&aeroport.CountryId,
			&aeroport.CityId,
			&aeroport.Latitude,
			&aeroport.Longitude,
			&aeroport.Radius,
			&aeroport.Image,
			&aeroport.Address,
			&aeroport.Country,
			&aeroport.City,
			&aeroport.SearchText,
			&aeroport.Code,
			&aeroport.ProductCount,
			&aeroport.Gmt,
			&aeroport.CreatedAt,
			&aeroport.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Aeroports = append(resp.Aeroports, &aeroport)
	}

	return &resp, nil
}

func (r *aeroportRepo) Update(req *models.UpdateAeroport) (int64, error) {

	query := `
		UPDATE aeroport
			SET
					"title" = $2
					"country_id" = $3
					"city_id" = $4
					"latitude" = $5
					"longitude" = $6
					"radius" = $7
					"image" = $8
					"address" = $9
					"country" = $10
					"city" = $11
					"search_text" = $12
					"code" = $13
					"product_count" = $14
					"gmt" = $15
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Latitude,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Address,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
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

func (r *aeroportRepo) Delete(req *models.AeroportPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM aeroport WHERE id = $1", req.Id)
	return err
}
