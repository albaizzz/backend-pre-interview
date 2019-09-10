package services

import (
	"testing"

	repomock "github.com/entraktest/internal/repositories/mocks"
	"github.com/golang/mock/gomock"
)

func Test_Store(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	deviceRepository := repomock.NewDeviceRepositoryMock(mockCtrl)
	svc := NewDeviceService(deviceRepository)
}

func Test_GetById(t *testing.T) {

}

func Test_GetAll(t *testing.T) {

}
