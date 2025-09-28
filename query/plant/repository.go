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
	query := "INSERT INTO plants (id, name, description, image_url, watering_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())"
	var wateringDate interface{}
	if plant.WateringDate != nil {
		wateringDate = plant.WateringDate.Format("2006-01-02")
	}
	
	_, err := r.db.Exec(query, plant.ID, plant.Name, plant.Description, plant.ImageURL, wateringDate)
	if err != nil {
		return fmt.Errorf("failed to create plant: %v", err)
	}
	
	return nil
}

func (r *Repository) save(plant entities.Plant) error {
	query := "UPDATE plants SET name = ?, description = ?, image_url = ?, watering_date = ?, updated_at = NOW() WHERE id = ?"
	var wateringDate interface{}
	if plant.WateringDate != nil {
		wateringDate = plant.WateringDate.Format("2006-01-02")
	}
	
	_, err := r.db.Exec(query, plant.Name, plant.Description, plant.ImageURL, wateringDate, plant.ID)
	if err != nil {
		return fmt.Errorf("failed to update plant: %v", err)
	}
	return nil
}

func (r *Repository) findByID(id string) (entities.Plant, error) {
	// DBからの取得
	return *entities.NewPlant("test"), nil
}

func (r *Repository) FindAll(limit int, offset int) ([]entities.Plant, error) {
	query := "SELECT id, name, description, image_url, watering_date, created_at, updated_at FROM plants LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("Error closing rows: %v\n", err)
		}
	}()

	var plants []entities.Plant
	for rows.Next() {
		var id string
		var name string
		var description sql.NullString
		var imageURL sql.NullString
		var wateringDate sql.NullString
		var createdAt time.Time
		var updatedAt time.Time
		err := rows.Scan(&id, &name, &description, &imageURL, &wateringDate, &createdAt, &updatedAt)
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

		var desc *string
		if description.Valid {
			desc = &description.String
		}

		var imgURL *string
		if imageURL.Valid {
			imgURL = &imageURL.String
		}

		plant := entities.Plant{
			ID:          id,
			Name:        plantName,
			Description: desc,
			ImageURL:    imgURL,
			WateringDate: wateringTime,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
		}
		plants = append(plants, plant)
	}

	return plants, nil
}

func (r *Repository) CreateWateringRecord(record entities.WateringRecord) error {
	query := "INSERT INTO watering_records (id, plant_id, watered_at, notes, created_at) VALUES (?, ?, ?, ?, NOW())"
	_, err := r.db.Exec(query, record.ID, record.PlantID, record.WateredAt, record.Notes)
	if err != nil {
		return fmt.Errorf("failed to create watering record: %v", err)
	}
	return nil
}

func (r *Repository) FindWateringRecordsByPlantID(plantID string) ([]entities.WateringRecord, error) {
	query := "SELECT id, plant_id, watered_at, notes, created_at FROM watering_records WHERE plant_id = ? ORDER BY watered_at DESC"
	rows, err := r.db.Query(query, plantID)
	if err != nil {
		return nil, fmt.Errorf("failed to query watering records: %v", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("Error closing rows: %v\n", err)
		}
	}()

	var records []entities.WateringRecord
	for rows.Next() {
		var record entities.WateringRecord
		var notes sql.NullString
		
		err := rows.Scan(&record.ID, &record.PlantID, &record.WateredAt, &notes, &record.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan watering record: %v", err)
		}
		
		if notes.Valid {
			record.Notes = &notes.String
		}
		
		records = append(records, record)
	}
	
	return records, nil
}
