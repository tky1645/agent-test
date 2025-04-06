package user

import (
	"DDD/entities"
	"database/sql"
	"fmt"

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
	fmt.Println("db connect success")
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user entities.User) error {
	query := "INSERT INTO users (id, name) VALUES (?, ?)"
	_, err := r.db.Exec(query, user.ID, user.Name)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}
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

func (r *UserRepository) GetAll() ([]entities.User, error) {
	query := "SELECT id, name FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entities.User{}
	for rows.Next() {
		var userData userTable
		err := rows.Scan(&userData.ID, &userData.Name)
		if err != nil {
			return nil, err
		}
		user, err := entities.NewUser(userData.ID, userData.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// get メソッドを fetchUserData に名称変更
func (r *UserRepository) fetchUserData(id int) userTable {
    // 仮のデータ取得処理
	query := "SELECT id, name FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var userID int
	var userName string
	err := row.Scan(&userID, &userName)
	if err != nil {
		fmt.Println("fetchUserData query error", err)
		return userTable{}
	}

	return userTable{
		ID:   userID,
		Name: userName,
	}
}

func (r *UserRepository) GetByID(id string) (entities.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var userId int
	var userName string
	var userEmail string
	var userPassword string
	err := row.Scan(&userId, &userName, &userEmail, &userPassword)
	if err != nil {
		fmt.Println("getByID query error", err)
		return entities.User{}, err
	}

	user, err := entities.NewUser(userId, userName)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Update(id string, name string) error {
	query := "UPDATE users SET name = ? WHERE id = ?"
	_, err := r.db.Exec(query, name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id uint) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
