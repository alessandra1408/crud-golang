package models

import "github.com/alessandra1408/crud-golang/db"

func GetAll() (todos []Todo, err error) {
	conn, cErr := db.NewPostgresConnection()
	if cErr != nil {
		return
	}
	defer conn.Close()

	rows, qErr := conn.Query(`SELECT * FROM todos`)
	if qErr != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, todo.Description, &todo.Done)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
