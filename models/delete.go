package models

import "github.com/alessandra1408/crud-golang/db"

func Delete(id int64) (int64, error) {
	conn, cErr := db.NewPostgresConnection()
	if cErr != nil {
		return 0, cErr
	}
	defer conn.Close()

	res, eErr := conn.Exec("DELETE from todos WHERE id=$1", id)
	if eErr != nil {
		return 0, eErr
	}

	return res.RowsAffected()
}
