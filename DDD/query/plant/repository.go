package plant

import (
	"DDD/entities"
	"database/sql"
)

var _ IPlantRepository  = (*Repository)(nil)

type Repository struct {
	db *sql.DB
}

func newRepo()*Repository{
	sql.Open("mysql", "user=sampleuser password= samplepass")
	return &Repository{
		db: nil,
	}
}
func (r *Repository) create(entities.Plant) error {

	return nil
}

func (r *Repository) save(entities.Plant) error {
	return nil
}

func (r *Repository) findByID(id int)(entities.Plant, error) {
	// DBからの取得
	return *entities.NewPlant("test"), nil
}