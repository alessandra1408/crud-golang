package models

import "github.com/alessandra1408/crud-golang/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.NewPostgresConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	err = row.Scan(&todo.ID, &todo.Title, todo.Description, &todo.Done)

	return
}
