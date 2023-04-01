package models

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCheckOwnerShip(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Unable to initialize mock DB", err)
	}

	// uses "gorm.io/driver/postgres" library
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	DB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal("Unable to initialize GORM", err)
	}

	mockRows := sqlmock.NewRows([]string{"id", "user_id"}).AddRow(1, 1)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "user_id" FROM "events" WHERE id = $1 AND "events"."deleted_at" IS NULL`)).WillReturnRows(mockRows)
	// EVENTs
	if err := checkOwnership[Event]("event", DB.Set(GORM_CONTEXT_USER_KEY, 1).Set("event", 1)); err != nil {
		t.Fatal("Check ownership failed, wanted err = nil", err)
	}

}
