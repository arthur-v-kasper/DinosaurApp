package dinosaur_test

import (
	"database/sql"
	"testing"

	"github.com/arthurkasper/DinosaurApp/core/dinosaur"
	_ "github.com/mattn/go-sqlite3"
)

func TestCRUD(t *testing.T) {
	d := &dinosaur.Dinosaur{
		ID:             1,
		Name:           "T-Rex",
		Era:            dinosaur.Jurassic,
		Classification: dinosaur.Theropods,
	}

	db, err := sql.Open("sqlite3", "../../data/dinosaur_test.db")
	if err != nil {
		t.Fatalf("Can't connect to data base... %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Can't clean the data base... %s", err.Error())
	}
	defer db.Close()

	service := dinosaur.NewService(db)

	t.Run("Store the Dinosaur", func(t *testing.T) {

		err = service.Store(d)
		if err != nil {
			t.Fatalf("It can't insert to data base.... %s", err)
		}
	})

	t.Run("Get the Dinosaur", func(t *testing.T) {

		saved, err := service.Get(1)

		if err != nil {
			t.Fatalf("It can't get to data base.... %s", err)
		}

		if saved.ID != 1 {
			t.Fatalf("Ivalid data, want %d got %d", 1, saved.ID)
		}
	})

}

func clearDB(d *sql.DB) error {
	tx, err := d.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from dinosaur")
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
