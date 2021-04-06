package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurkasper/DinosaurApp/core/dinosaur"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type DinosaurOperationServiceMock struct{}

func (d DinosaurOperationServiceMock) GetAll() ([]*dinosaur.Dinosaur, error) {
	dino1 := &dinosaur.Dinosaur{
		ID:             1,
		Name:           "T-Rex",
		Era:            dinosaur.Jurassic,
		Classification: dinosaur.Theropods,
	}

	dino2 := &dinosaur.Dinosaur{
		ID:             2,
		Name:           "Triceratops",
		Era:            dinosaur.Jurassic,
		Classification: dinosaur.Theropods,
	}

	return []*dinosaur.Dinosaur{dino1, dino2}, nil
}

func (d DinosaurOperationServiceMock) Get(ID int64) (*dinosaur.Dinosaur, error) {
	dino1 := &dinosaur.Dinosaur{
		ID:             ID,
		Name:           "T-Rex",
		Era:            dinosaur.Jurassic,
		Classification: dinosaur.Theropods,
	}

	return dino1, nil
}

func (d DinosaurOperationServiceMock) Store(*dinosaur.Dinosaur) error {
	return nil
}

func (d DinosaurOperationServiceMock) Update(*dinosaur.Dinosaur) error {
	return nil
}

func (d DinosaurOperationServiceMock) Remove(int64) error {
	return nil
}

func TestGetAll(t *testing.T) {
	service := &DinosaurOperationServiceMock{}
	handler := getAllDinosaur(service)
	r := mux.NewRouter()
	r.Handle("/v1/dinosaur", handler)

	request, err := http.NewRequest("GET", "/v1/dinosaur", nil)
	assert.Nil(t, err)
	request.Header.Set("Content-type", "json/application")

	responseRequest := httptest.NewRecorder()
	r.ServeHTTP(responseRequest, request)
	assert.Equal(t, http.StatusOK, responseRequest.Code)

	var result []*dinosaur.Dinosaur
	err = json.NewDecoder(responseRequest.Body).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, int64(2), result[1].ID)
}

func TestGetAllClassification(t *testing.T) {
	dc, _ := dinosaur.GetAllClassification()

	assert.Equal(t, len(dc), 6)
}
