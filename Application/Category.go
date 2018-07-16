
package application

import (
	"awesomeProject2/config"
	"awesomeProject2/Domain"
	_ "awesomeProject2/Infrastructure/persistance"
	"awesomeProject2/Infrastructure/persistance"
)


func AddCategory(Category *Domain.Category) error {
	conn, err := config.NewDBConnection()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app1
	}
	defer conn.Close()
	repository := persistence.NewCategoryRepositoryWithRDB(conn)

	return repository.Add(Category)
}