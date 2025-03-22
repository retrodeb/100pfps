package model

import "github.com/retrodeb/100pfps/db"

func SetupMigrations() {
	db.GetDB().AutoMigrate(
		&Profile{},
		&Tag{},
		&Ban{},
		&Admin{},
	)
}
