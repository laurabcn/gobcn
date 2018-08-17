package persistence

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // driver
	categoryDomain "github.com/laurabcn/gobcn/Domain"
)

type CategorySQLRepository struct {
	Conn *sql.DB
}
func NewCategoryRepositoryWithRDB(conn *sql.DB) categoryDomain.CategoryRepository {
	return &CategorySQLRepository{Conn: conn}
}

func (r *CategorySQLRepository) Add(Category *categoryDomain.Category) error {
	stmtIns, err := r.Conn.Prepare("INSERT INTO categories VALUES (?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	if err != nil {
		panic("That's embarrassing...")
	}

	_, err = stmtIns.Exec(Category.Id, Category.Name, Category.Language, Category.Mostra, Category.Position)
	if err != nil {
		panic(err.Error())
	}
	return err
}

