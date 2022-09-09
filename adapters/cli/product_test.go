package cli_test

import (
	"fmt"
	"hexagonal/adapters/cli"
	mock_application "hexagonal/application/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Save(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectResult := fmt.Sprintf("Product with Id: %s, name: %s, price: %f, and status: %s has been created",
		productMock.GetId(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectResult, result)

	expectResult = fmt.Sprintf("Product: %s been enabled", productMock.GetName())
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectResult, result)

	expectResult = fmt.Sprintf("Product: %s been disabled", productMock.GetName())
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectResult, result)

	expectResult = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productMock.GetId(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectResult, result)
}
