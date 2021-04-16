package dinosaur

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var d = &Dinosaur{
	ID:             1,
	Name:           "T-Rex",
	Era:            Jurassic,
	Classification: Theropods,
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

func TestCRUD(t *testing.T) {

	db, err := sql.Open("sqlite3", "../../data/dinosaur_test.db")
	if err != nil {
		t.Fatalf("Can't connect to data base... %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Can't clean the data base... %s", err.Error())
	}
	defer db.Close()

	service := NewService(db)

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

//test using mock
func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurrer when opening a mock database connect: '%s'", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("update dinosaur").
		ExpectExec().
		WithArgs("T-Rex", int64(2), int64(1), int64(1)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mockDB := NewService(db)

	if err = mockDB.Update(d); err != nil {
		t.Errorf("error was not expected while updating dinosaur: %s", err)
	}

}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurrer when opening a mock database connect: '%s'", err)
	}

	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("delete from dinosaur").
		WithArgs(int64(1)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mockDB := NewService(db)

	if err = mockDB.Remove(d.ID); err != nil {
		t.Errorf("error was not expected while remove the dinosaur: %s", err)
	}

}

func TestGetClassification(t *testing.T) {
	dc, err := GetAllClassification()
	if err != nil {
		t.Errorf("error was not expected: %s", err)
	}

	assert.Contains(t, dc[0].Name, "Theropods")
	assert.Equal(t, len(dc), 5)

}
