package application

type DatabaseDatasource interface {
	Save(collection string, data interface{}) (error error, value map[string]string)
	Delete(collection string, id string) error
	DeleteWhere(collection string, where QueryFilter) error
	Select(collection string, where QueryFilter) (error, interface{})
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
