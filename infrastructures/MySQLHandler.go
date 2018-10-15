package infrastructures

import (
	"database/sql"
	"fmt"

	"github.com/VanaraID/clean-architecture/interfaces"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLHandler struct {
	Conn *sql.DB
}

func (handler *MySQLHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MySQLHandler) Query(statement string) (interfaces.IRow, error) {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(MySQLRow), err
	}
	row := new(MySQLRow)
	row.Rows = rows

	return row, nil
}

type MySQLRow struct {
	Rows *sql.Rows
}

func (r MySQLRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r MySQLRow) Next() bool {
	return r.Rows.Next()
}
