package usecases

import (
	"database/sql"
	"log"

	"github.com/ilhammhdd/kudaki-user-service/entities"

	"github.com/ilhammhdd/go_error"
)

type DBOperator interface {
	Command() error
	Query() (*sql.Rows, error)
	QueryRow() (*sql.Row, error)
	QueryRowsToMap() (*[]map[string]interface{}, error)
}

type DBOperation struct {
	Stmt string
	Args []interface{}
}

func (dbo *DBOperation) Command() error {
	log.Println("DBO COMMAND UNDERWAY : ", *dbo)
	outStmt, err := entities.DB.Prepare(dbo.Stmt)
	defer outStmt.Close()
	if err != nil {
		return err
	}

	_, err = outStmt.Exec(dbo.Args...)
	if err != nil {
		return err
	}

	return nil
}

func (dbo *DBOperation) Query() (*sql.Rows, error) {
	log.Println("DBO QUERY UNDERWAY")
	outStmt, err := entities.DB.Prepare(dbo.Stmt)
	defer outStmt.Close()
	go_error.HandleError(err)

	rows, err := outStmt.Query(dbo.Args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (dbo *DBOperation) QueryRow() (*sql.Row, error) {
	outStmt, err := entities.DB.Prepare(dbo.Stmt)
	defer outStmt.Close()
	if err != nil {
		return nil, err
	}

	resultRow := outStmt.QueryRow(dbo.Args...)

	return resultRow, nil
}

func (dbo *DBOperation) QueryRowsToMap() (*[]map[string]interface{}, error) {
	outStmt, err := entities.DB.Prepare(dbo.Stmt)
	defer outStmt.Close()
	if err != nil {
		return nil, err
	}

	rows, err := outStmt.Query(dbo.Args...)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))

	var resultRows []map[string]interface{}
	resultRow := make(map[string]interface{})

	for rows.Next() {
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		err = rows.Scan(columnPointers...)
		if err != nil {
			return nil, err
		}

		columnTypes, _ := rows.ColumnTypes()

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			if (*columnTypes[i]).DatabaseTypeName() == "TINYINT" {
				if (*val).(int64) == 1 {
					resultRow[colName] = true
				} else if (*val).(int64) == 0 {
					resultRow[colName] = false
				}
			} else {
				resultRow[colName] = *val
			}
		}

		resultRows = append(resultRows, resultRow)
	}

	return &resultRows, nil
}

func DBCommandOperation(dbo DBOperator) {
	dbo.Command()
}

func DBQueryOperation(dbo DBOperator) (*sql.Rows, error) {
	return dbo.Query()
}

func DBQueryRowsToMap(dbo DBOperator) (*[]map[string]interface{}, error) {
	return dbo.QueryRowsToMap()
}
