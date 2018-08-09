
package application

import (
	"github.com/laurabcn/gobcn/config"
	categoryDomain "github.com/laurabcn/gobcn/Domain"
	"github.com/laurabcn/gobcn/Infrastructure/persistance"
)


func AddCategory(Category *categoryDomain.Category) error {
	conn, err := config.NewDBConnection()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app1
	}
	defer conn.Close()
	repository := persistence.NewCategoryRepositoryWithRDB(conn)
}