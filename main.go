package main

import (
	"github.com/aarondl/sqlboiler-sqlite3/driver"
	"github.com/aarondl/sqlboiler/v4/drivers"
)

func main() {
	drivers.DriverMain(&driver.SQLiteDriver{})
}
