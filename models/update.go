package models

import "github.com/alessandra1408/crud-golang/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, cErr := db.NewPostgresConnection()
	if cErr != nil {
		return 0, cErr
	}
	defer conn.Close()

	res, eErr := conn.Exec("UPDATE todos SET id=$1, title=$2, description=$3, done=$4", todo.ID, todo.Title, todo.Description, todo.Done)
	if eErr != nil {
		return 0, eErr
	}

	return res.RowsAffected()
}
