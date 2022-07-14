package controller

import(
	"encoding/json"
	"net/http"

	"github.com/VladRomanciuc/Go-classes/api/models"
)

var carDetailsService models.CarDetailsService

type carcontroller struct{}

func NewCarDetailsController(service models.CarDetailsService) models.CarDetailsController{
	carDetailsService = service
	return &carcontroller{}
}

func(* carcontroller) GetCarDetails(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	result := carDetailsService.GetDetails()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}