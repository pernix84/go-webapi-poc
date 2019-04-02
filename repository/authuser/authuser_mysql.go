package authuserrepository

import (
	"database/sql"
	"go-webapi-poc/models"
	"log"
)

type AuthUserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (a AuthUserRepository) GetUsers(db *sql.DB, user models.AuthUser, users []models.AuthUser) ([]models.AuthUser, error) {
	rows, err := db.Query("select * from oauth_user")

	if err != nil {
		return []models.AuthUser{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FullName, &user.LoginCount, &user.OAuthId)
		users = append(users, user)
	}

	if err != nil {
		return []models.AuthUser{}, err
	}

	return users, nil
}

func (a AuthUserRepository) GetUser(db *sql.DB, user models.AuthUser, id int) (models.AuthUser, error) {
	rows := db.QueryRow("select * from oauth_user where oauth_user_id = ?", id)

	err := rows.Scan(&user.ID, &user.FullName, &user.LoginCount, &user.OAuthId)
	if err != nil {
		return models.AuthUser{}, err
	}

	return user, nil
}

func (a AuthUserRepository) AddUser(db *sql.DB, user models.AuthUser) (int, error) {
	res, err := db.Exec("insert into oauth_user (full_name, login_count, oauth_id) values(?, ?, ?)",
		user.FullName, user.LoginCount, user.OAuthId)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (a AuthUserRepository) UpdateUser(db *sql.DB, user models.AuthUser) (int64, error) {
	result, err := db.Exec("update oauth_user set full_name=?, login_count=?, oauth_id=? where oauth_user_id=?",
		&user.FullName, &user.LoginCount, &user.OAuthId, &user.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (a AuthUserRepository) RemoveUser(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from oauth_user where oauth_user_id = ?", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
