package infra

import (
	"auth_api/config/database"
	"auth_api/core/application"
)

type GoOrmDatasource struct {
	Database database.GoOrmDatabase
}

func(this *GoOrmDatasource)	Save(collection string, data map[string]interface{}) application.QueryResult{
	result := this.Database.Create(&data)
	if(result.Error !=nil){
		return application.QueryResult{
			Failed: true,
			Errors: error[result.Error],
			Data:   nil,
		}
	}

	return application.QueryResult{
		Failed: false,
		Errors: nil,
		Data:   interface{}(data),
	}
}
func(this *GoOrmDatasource)Delete(data interface{}, id string) QueryResult{
	result := this.Database.Where("id = ?", id).Delete(&data)
	if(result.Error !=nil){
		return application.QueryResult{
			Failed: true,
			Errors: error[result.Error],
			Data:   nil,
		}
	}

	return application.QueryResult{
		Failed: false,
		Errors: nil,
		Data:   interface{}(data),
	}
}
	DeleteWhere(collection string, where QueryFilter) QueryResult
	Select(collection string, where QueryFilter) QueryResult