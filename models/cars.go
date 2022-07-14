package models

import (
	"net/http"
)

type CarDetailsController interface {
	GetCarDetails(w http.ResponseWriter, r *http.Request)
}

type CarDetailsService interface {
	GetDetails() CarDetails
}

type CarService interface {
	FetchData()
}
type OwnerService interface {
	FetchData()
}

type Car struct{
	CarData `json:"Car"`
}

type CarData struct{
	Id int `json:"id"`
	Brand string `json:"car"`
	Model string `json:"car_model"`
	Year int `json:"car_model_year"`
}

type Owner struct {
	OwnerData `json:"User"`
}

type OwnerData struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}

type CarDetails struct{
	Id int `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year int `json:"model_year"`
	FirstName string `json:"owner_first_name"`
	LastName string `json:"owner_last_name"`
	Email string `json:"owner_email"`
}

