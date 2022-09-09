package application_test

import (
	"hexagonal/application"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "Hello",
		Status: application.DISABLED,
		Price:  10,
	}

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "the price must be greater than zaro to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		Name:   "Hello",
		Status: application.DISABLED,
		Price:  10,
	}

	err := product.Disable()
	require.Equal(t, "the price must be equal to zaro to disable the product", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		Id:     uuid.NewV4().String(),
		Name:   "Hello",
		Status: application.DISABLED,
		Price:  10,
	}

	isValid, _ := product.IsValid()
	require.True(t, isValid)

	product.Status = "INVALID"

	isValid, err := product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the status must be enable or disable", err.Error())

	product.Status = application.DISABLED
	product.Price = -10

	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the price must be greater than or equal to zero", err.Error())
}
