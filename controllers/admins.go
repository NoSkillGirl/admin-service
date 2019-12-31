package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NoSkillGirl/admin-service/models"
	userController "github.com/NoSkillGirl/user-service/controllers"
)

//NewCompanyRequest struct
type NewCompanyRequest struct {
	Name      string
	OwnerName string
	PhoneNo   string
}

//NewBusRequest struct
type NewBusRequest struct {
	Number        string
	AC            bool `json:"ac"`
	Sleeper       bool
	TotalSeats    int
	Source        string
	Destination   string
	CompanyID     int
	DepartureTime string
	ArrivalTime   string
}

//NewCompany function
func NewCompany(w http.ResponseWriter, r *http.Request) {
	var reqJSON NewCompanyRequest
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		panic(err)
	}
	log.Println(reqJSON)

	errOccured, duplicateExist := models.AddCompany(reqJSON.Name, reqJSON.OwnerName, reqJSON.PhoneNo)
	if errOccured == true {
		data := userController.Response{
			Status:   500,
			Response: userController.ResponseMsg{},
			Error: userController.ErrorMessage{
				Msg: "Internal server Error",
			},
		}
		json.NewEncoder(w).Encode(data)
	} else if errOccured == false && duplicateExist == true {
		data := userController.Response{
			Status:   500,
			Response: userController.ResponseMsg{},
			Error: userController.ErrorMessage{
				Msg: "Company already exist",
			},
		}
		json.NewEncoder(w).Encode(data)
	} else {
		data := userController.Response{
			Status: 200,
			Response: userController.ResponseMsg{
				Msg: "Company succesfully created",
			},
			Error: userController.ErrorMessage{},
		}
		json.NewEncoder(w).Encode(data)
	}
}

//NewBus function
func NewBus(w http.ResponseWriter, r *http.Request) {
	var reqJSON NewBusRequest
	fmt.Println(r)
	err := json.NewDecoder(r.Body).Decode(&reqJSON)

	if err != nil {
		panic(err)
	}
	log.Println(reqJSON)

	errOccured, duplicateExist := models.AddBus(reqJSON.Number, reqJSON.AC, reqJSON.Sleeper, reqJSON.TotalSeats, reqJSON.Source, reqJSON.Destination, reqJSON.CompanyID, reqJSON.DepartureTime, reqJSON.ArrivalTime)
	if errOccured == true {
		data := userController.Response{
			Status:   500,
			Response: userController.ResponseMsg{},
			Error: userController.ErrorMessage{
				Msg: "Internal server Error",
			},
		}
		json.NewEncoder(w).Encode(data)
	} else if errOccured == false && duplicateExist == true {
		data := userController.Response{
			Status:   500,
			Response: userController.ResponseMsg{},
			Error: userController.ErrorMessage{
				Msg: "Bus already exist",
			},
		}
		json.NewEncoder(w).Encode(data)
	} else {
		data := userController.Response{
			Status: 200,
			Response: userController.ResponseMsg{
				Msg: "Bus succesfully added",
			},
			Error: userController.ErrorMessage{},
		}
		json.NewEncoder(w).Encode(data)
	}
}
