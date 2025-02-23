package plant

import "DDD/entities"

var _ IPlantRepository  = (*Repository)(nil)

type Repository struct {}

func newRepo()*Repository{
	return &Repository{}
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