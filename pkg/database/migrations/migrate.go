package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/leonardonatali/file-metadata-api/pkg/config"
)

func Migrate(cfg *config.Config) {
	// wrap assets into Resource
	s := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, cfg.GetDatabaseDSN(true))
	if err != nil {
		panic(err)
	}

	m.Up()
}
