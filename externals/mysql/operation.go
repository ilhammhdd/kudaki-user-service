package mysql

import (
	"database/sql"

	"github.com/ilhammhdd/go_tool/go_error"
)

type DBOperation struct{}

func NewDBOperation() DBOperation {
	return DBOperation{}
}

func (dbo DBOperation) Command(stmt string, args ...interface{}) error {
	outStmt, err := DB.Prepare(stmt)
	defer outStmt.Close()
	if err != nil {
		return err
	}

	_, err = outStmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func (dbo DBOperation) Query(stmt string, args ...interface{}) (*sql.Rows, error) {
	outStmt, err := DB.Prepare(stmt)
	defer outStmt.Close()
	go_error.ErrorHandled(err)

	rows, err := outStmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (dbo DBOperation) QueryRow(stmt string, args ...interface{}) (*sql.Row, error) {
	outStmt, err := DB.Prepare(stmt)
	defer outStmt.Close()
	if err != nil {
		return nil, err
	}

	resultRow := outStmt.QueryRow(args...)

	return resultRow, nil
}

func (dbo DBOperation) QueryRowsToMap(stmt string, args ...interface{}) (*[]map[string]interface{}, error) {
	outStmt, err := DB.Prepare(stmt)
	defer outStmt.Close()
	if err != nil {
		return nil, err
	}

	rows, err := outStmt.Query(args...)
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