package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VladRomanciuc/Go-classes/api/models"
)

const (
	carUrl = "https://myfakeapi.com/api/cars/1"
	userUrl = "https://myfakeapi.com/api/users/1"
)

type fetchCarDataService struct{}
func NewCarService() models.CarService{
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData(){
	client := http.Client{}
	response, _ := client.Get(carUrl)
	carChannel <-response
}

type fetchOwnerDataService struct{}
func NewOwnerCarService() models.CarService{
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData(){
	client := http.Client{}
	response, _ := client.Get(userUrl)
	ownerChannel <-response
}


var (
	carService models.CarService = NewCarService()
	ownerService models.OwnerService = NewOwnerCarService()
	carChannel = make(chan *http.Response)
	ownerChannel = make(chan *http.Response)
)

type carservice struct{}

func NewCarDetailsService() models.CarDetailsService{
	return &carservice{}
}


func(*carservice) GetDetails() models.CarDetails {
	//build 2 gorutines for https://myfakeapi.com/api/cars/1 and https://myfakeapi.com/api/users/1
	go carService.FetchData()
	go ownerService.FetchData()
	//create 2 channels to get the data from fake api
	car, _ := getCarData()
	owner, _ := getOwnerData()

	//assign the recieved values
	return models.CarDetails{
		Id: car.Id,
		Brand: car.Brand,
		Model: car.Model,
		Year: car.Year,
		FirstName: owner.FirstName,
		LastName: owner.LastName,
		Email: owner.Email,
	}
}

func getCarData() (models.Car, error) {
	r := <-carChannel
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil{
		log.Fatal(err)
		return car, err
	}
	return car, nil
}

func getOwnerData() (models.Owner, error){
	r := <-ownerChannel
	var owner models.Owner
	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil{
		log.Fatal(err)
		return owner, err
	}
	return owner, nil
}