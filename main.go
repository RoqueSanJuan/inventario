package main

import (
	"context"

	"github.com/RoqueSanJuan/go-inventario/database"
	"github.com/RoqueSanJuan/go-inventario/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, error := db.Query("select * from USERS")
				if error != nil {
					panic(error)
				}
			},
		),
	)

	app.Run()
}
