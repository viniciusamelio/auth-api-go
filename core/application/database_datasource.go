package application

import "auth_api/config/database"

type DatabaseDatasource interface {
	Save(data database.Schema) QueryResult[database.Schema]
	Delete(data database.Schema, id string) QueryResult[database.Schema]
	Select(data database.Schema, where string, args ...interface{}) QueryResult[database.Schema]
}

type QueryResult[T database.Schema] struct {
	Failed bool
	Errors []error
	Data   []T
}
