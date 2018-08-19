package persistence

import "database/sql"

type SiteCatSQLRepository struct {
	Conn *sql.DB
}
func NewSiteCatRepositoryWithRDB(conn *sql.DB) Domain.SiteRepository {
	return &SiteCatSQLRepository{Conn: conn}
}

func (r *SiteCatSQLRepository) AddSite(Site *Domain.Site, Category *Domain.Category) error {
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
