package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	carDetailsService = NewCarDetailsService()
)

func TestGetDetails(t *testing.T) {
	carDetails := carDetailsService.GetDetails()

	assert.NotNil(t, carDetails)
	assert.Equal(t, 1, carDetails.Id)
	assert.Equal(t, "Montero", carDetails.Model)
	assert.Equal(t, "Caldero", carDetails.LastName)
}