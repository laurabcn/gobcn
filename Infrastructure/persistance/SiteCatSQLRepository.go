package persistence

import (
	"database/sql"
	"github.com/laurabcn/gobcn/Domain"
	"github.com/satori/go.uuid"
)
type SiteCatSQLRepository struct {
	Conn *sql.DB
}
func NewSiteCatRepositoryWithRDB(conn *sql.DB) Domain.SiteCatRepository {
	return &SiteCatSQLRepository{Conn: conn}
}

func (r *SiteCatSQLRepository) AddSiteCategory(Site *Domain.Site, Category *Domain.Category) error {
	stmtIns, err := r.Conn.Prepare("INSERT INTO siteCategories VALUES (?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	if err != nil {
		panic("That's embarrassing...")
	}

	_, err = stmtIns.Exec(
		uuid.Must(uuid.NewV4()),
		Site.Id,
		Category.Id,
	)
	if err != nil {
		panic(err.Error())
	}
	return err
}
