package Application

import (
	"github.com/laurabcn/gobcn/config"
	"github.com/laurabcn/gobcn/Domain"
	"github.com/laurabcn/gobcn/Infrastructure/persistance"
)


func AddSite(Site *Domain.Site) error {
	conn, err := config.NewDBConnection()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app1
	}
	defer conn.Close()
	repository := persistence.NewSiteRepositoryWithRDB(conn)

	return repository.AddSite(Site)
}
