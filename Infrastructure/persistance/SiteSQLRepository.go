package persistence

import (
	_ "github.com/go-sql-driver/mysql" // driver
	"github.com/laurabcn/gobcn/Domain/Site"
	"database/sql"
	"github.com/laurabcn/gobcn/Domain"
)

type SiteSQLRepository struct {
	Conn *sql.DB
}
func NewSiteRepositoryWithRDB(conn *sql.DB) Domain.SiteRepository {
	return &SiteSQLRepository{Conn: conn}
}

func (r *SiteSQLRepository) AddSite(Site *Domain.Site) error {
	stmtIns, err := r.Conn.Prepare("INSERT INTO sites VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	if err != nil {
		panic("That's embarrassing...")
	}

	_, err = stmtIns.Exec(
		Site.Id,
		Site.Name,
		Site.Language,
		Site.Mostra,
		Site.District,
		Site.Phone,
		Site.Web,
		Site.Content,
		Site.Excerpt,
		Site.Latitude,
		Site.Longitude,
		Site.Type,
		Site.Barri,
		Site.Address)
	if err != nil {
		panic(err.Error())
	}
	return err
}
