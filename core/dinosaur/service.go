package dinosaur

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type OperationService interface {
	GetAll() ([]*Dinosaur, error)
	Get(ID int64) (*Dinosaur, error)
	Store(d *Dinosaur) error
	Update(d *Dinosaur) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Dinosaur, error) {
	var result []*Dinosaur

	rows, err := s.DB.Query("select id, name, era, classification from dinosaur")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	// Append the result
	for rows.Next() {
		var d Dinosaur

		err = rows.Scan(&d.ID, &d.Name, &d.Era, &d.Classification)
		if err != nil {
			return nil, err
		}
		result = append(result, &d)
	}

	return result, nil
}

func (s *Service) Get(ID int64) (*Dinosaur, error) {
	var d Dinosaur

	stmt, err := s.DB.Prepare("select id, name, era, classification from dinosaur where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(ID).Scan(&d.ID, &d.Name, &d.Era, &d.Classification)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (s *Service) Store(d *Dinosaur) error {
	insert, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := insert.Prepare("insert into dinosaur(id, name, era, classification) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.ID, d.Name, d.Era, d.Classification)
	if err != nil {
		insert.Rollback()
		return err
	}
	insert.Commit()
	return nil
}

func (s *Service) Update(d *Dinosaur) error {
	if d.ID == 0 {
		return fmt.Errorf("Invalid ID")
	}

	update, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := update.Prepare("update dinosaur set name=?, era=?, classification=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.Name, d.Era, d.Classification, d.ID)
	if err != nil {
		update.Rollback()
		return err
	}

	update.Commit()
	return nil
}

func (s *Service) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("Invalid ID")
	}

	delete, err := s.DB.Begin()
	if err != nil {
		return err
	}

	_, err = delete.Exec("delete from dinosaur where id=?", ID)
	if err != nil {
		delete.Rollback()
		return err
	}

	delete.Commit()
	return nil
}

func GetAllClassification() (map[DinosaurClassification]string, error) {
	return DinosaurClassificationMap, nil
}
