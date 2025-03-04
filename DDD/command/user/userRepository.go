package user

import (
	"DDD/entities"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var _ IUserRepository = (*UserRepository)(nil)

type userTable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user entities.User) error {
	return nil
}

// Create の戻り値に error を追加してエラーを呼び出し元へ伝播する
func (r *UserRepository) Create(id int) (entities.User, error) {
	userData := r.fetchUserData(id)
	user, err := entities.NewUser(userData.ID, userData.Name)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

// get メソッドを fetchUserData に名称変更
func (r *UserRepository) fetchUserData(id int) userTable {
    // 仮のデータ取得処理
	return userTable{
		ID:   id,
		Name: "getJohn",
	}
}

func (r *UserRepository) GetByID(id string) (entities.User, error) {
	// TODO: Implement GetByID logic here
	// Replace with actual database query
	// Example:
	// query := "SELECT id, name FROM users WHERE id = ?"
	// row := r.db.QueryRow(query, id)
	return entities.NewUser(1, "getJohn")
}

func (r *UserRepository) Update(id string, name string) error {
	// TODO: Implement Update logic here
	// Replace with actual database update
	// Example:
	// query := "UPDATE users SET name = ? WHERE id = ?"
	// _, err := r.db.Exec(query, name, id)
	return nil
}

func (r *UserRepository) Delete(id uint) error {
	// TODO: Implement Delete logic here
	// Replace with actual database delete
	// Example:
	// query := "DELETE FROM users WHERE id = ?"
	// _, err := r.db.Exec(query, id)
	return nil
}
