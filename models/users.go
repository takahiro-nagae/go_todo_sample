package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmdU := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmdU, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmdU := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err = Db.Exec(cmdU, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) DeleteUser() (err error) {
	cmdU := `DELETE FROM users WHERE id = ?`
	_, err = Db.Exec(cmdU, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
