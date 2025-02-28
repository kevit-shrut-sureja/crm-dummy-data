package main

type database struct {
	DB string
}

func (d *database) GetDB() string {
	// return the database name
	return d.DB
}
