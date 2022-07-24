package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"practiceproject/database"
	_ "practiceproject/docs"
	"practiceproject/entity"
)

// CreateResource godoc
// @Summary Creates a account
// @Description creates Resource directory
// @Tags Actions with account
// @Param first-name path string true "name"
// @Param last-name path string true "uuid"
// @Param id path integer($int64) true "uuid"
// @Param user-name path string true "uuid"
// @Param user-password path string true "uuid"
// @Accept  json
// @Success 200  {object} string "successful"
// @Failure 400  {string} string "uncorrect"
// @Failure 404  {string} string "Error. Not found"
// @Failure 500  {string} string "system error"
// @Router /person/create [post]
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person entity.User
	json.NewDecoder(r.Body).Decode(&person)
	database.Instance.Create(&person)
	json.NewEncoder(w).Encode(person)

}
func checkIfExists(productId string) bool {
	var person entity.User
	database.Instance.First(&person, productId)
	if person.ID == 0 {
		return false
	}
	return true
}

// GetPerson godoc
// @Summary get persons by ID
// @Description creates Resource directory
// @Tags Get
// @Param ID path string true "uuid"
// @Accept  json
// @Success 200  {object} string "successful"
// @Failure 400  {string} string "uncorrect"
// @Failure 404  {string} string "Error. Not found"
// @Failure 500  {string} string "system error"
// @Router /person/getpersonbyid [get]
func GetPersonById(w http.ResponseWriter, r *http.Request) {
	personId := mux.Vars(r)["id"]
	if checkIfExists(personId) == false {
		json.NewEncoder(w).Encode("Person not found")
		return
	}
	var product entity.User
	database.Instance.First(&product, personId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

// GetPerson godoc
// @Summary get persons not by ID
// @Description creates Resource directory
// @Tags Get
// @Param first-name path string true "name"
// @Param last-name path string true "uuid"
// @Param id path integer($int64) true "uuid"
// @Param user-name path string true "uuid"
// @Param user-password path string true "uuid"
// @Accept  json
// @Success 200  {object} string "successful"
// @Failure 400  {string} string "uncorrect"
// @Failure 404  {string} string "Error. Not found"
// @Failure 500  {string} string "system error"
// @Router /person/getpersons[get]
func GetPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("dhfhdfsgs")
	persons := []entity.User{}
	database.Instance.Find(&persons)
	fmt.Printf("%#v", persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)

}

// UpdateAccount godoc
// @Summary updates account
// @Description creates Resource directory
// @Tags Actions with account
// @Param first-name path string true "name"
// @Param last-name path string true "uuid"
// @Param id path integer($int64) true "uuid"
// @Param user-name path string true "uuid"
// @Param user-password path string true "uuid"
// @Accept  json
// @Success 200  {object} string "successful"
// @Failure 400  {string} string "uncorrect"
// @Failure 404  {string} string "Error. Not found"
// @Failure 500  {string} string "system error"
// @Router /person/update [patch]
func Update(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfExists(productId) == false {
		json.NewEncoder(w).Encode("Person not found")
		return
	}
	var person entity.User
	database.Instance.First(&person, productId)
	json.NewDecoder(r.Body).Decode(&person)
	database.Instance.Save(&person)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)

}

// DeleteAccount godoc
// @Summary delete account
// @Description creates Resource directory
// @Tags Actions with account
// @Param first-name path string true "name"
// @Param last-name path string true "uuid"
// @Param id path integer($int64) true "uuid"
// @Param user-name path string true "uuid"
// @Param user-password path string true "uuid"
// @Accept  json
// @Success 200  {object} string "successful"
// @Failure 400  {string} string "uncorrect"
// @Failure 404  {string} string "Error. Not found"
// @Failure 500  {string} string "system error"
// @Router /person/delete [delete]
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personId := mux.Vars(r)["id"]
	if checkIfExists(personId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Person not found")
		return
	}
	var person entity.User
	database.Instance.Delete(&person, personId)
	json.NewEncoder(w).Encode("Person deleted ")
}
