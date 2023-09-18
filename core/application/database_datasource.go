package application

type DatabaseDatasource interface {
	Save(data map[string]interface{}) QueryResult
	Delete(data interface{}, id string) QueryResult
	DeleteWhere(data interface{}, where QueryFilter) QueryResult
	Select(collection string, where QueryFilter) QueryResult
}

type QueryResult struct {
	Failed bool
	Errors []error
	Data   []interface{}
}

type QueryFilter struct {
	Field            string
	Value            interface{}
	Operator         string
	AggregateFilters []AggregateFilters
}

type AggregateFilters struct {
	Field           string
	Operator        string
	Value           interface{}
	AggregationType AggregateType
}

type AggregateType string

const (
	And AggregateType = "AND"
	Or                = "OR"
)
