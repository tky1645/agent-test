package plant

import (
	"DDD/entities"
	"database/sql"
	"fmt"
	"time"
)

type Repository struct {
	db *sql.DB
}

func newRepo() *Repository {
	db, err := sql.Open("mysql", "sampleuser:samplepass@tcp(ddd_rdb:3306)/sampledb")
	if err != nil {
		fmt.Println("db err", err)
		panic(err)
	}
	return &Repository{
		db: db,
	}
}

func (r *Repository) create(plant entities.Plant) error {
	return nil
}

func (r *Repository) save(plant entities.Plant) error {
	return nil
}

func (r *Repository) findByID(id int) (entities.Plant, error) {
	// DBからの取得
	return *entities.NewPlant("test"), nil
}

func (r *Repository) FindAll(limit int, offset int) ([]entities.Plant, error) {
	query := "SELECT id, name, wateringDate, created_at, updated_at FROM plant LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plants []entities.Plant
	for rows.Next() {
		var id int
		var name string
		var wateringDate sql.NullString
		var createdAt time.Time
		var updatedAt time.Time
		err := rows.Scan(&id, &name, &wateringDate, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		plantName, err := entities.NewPlantName(name)
		if err != nil {
			fmt.Println("Error creating plant name:", err)
			continue // Skip this plant if there's an error
		}

		var wateringTime *time.Time
		if wateringDate.Valid {
			// Parse the wateringDate string to time.Time
			t, err := time.Parse("2006-01-02", wateringDate.String) // Adjust the format as needed
			if err != nil {
				fmt.Println("Error parsing watering date:", err)
				continue // Skip this plant if there's an error
			}
			wateringTime = &t
		}

		plant := entities.Plant{
			ID:   id,
			Name: plantName,
			WateringDate: wateringTime,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		}
		plants = append(plants, plant)
	}

	return plants, nil
}
