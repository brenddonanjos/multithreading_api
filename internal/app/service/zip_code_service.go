package service

import (
	"errors"
	"time"

	"github.com/brenddonanjos/multithreading_api/internal/entity"
	"github.com/brenddonanjos/multithreading_api/internal/infra/interfaces"
)

func GetZipCodeInfo(zipCode string) (*entity.ZipCode, error) {
	zipCopdeFinderInterfaces := []interfaces.ZipcodeFinderInterface{
		NewBrasilApiService(),
		NewViaCepService(),
	}
	startTime := time.Now()
	channel := make(chan *entity.ZipCode)

	for _, zipCodeFinderInterface := range zipCopdeFinderInterfaces {
		go publisher(zipCode, zipCodeFinderInterface, channel, startTime)

	}

	return subscriber(channel)
}

func publisher(zipCode string, zipCodeFinderInterface interfaces.ZipcodeFinderInterface, channel chan<- *entity.ZipCode, startTime time.Time) {
	zipCodeInfo, err := zipCodeFinderInterface.FetchZipCode(zipCode, startTime)
	if err != nil {
		panic(err.Error())
	}
	channel <- zipCodeInfo

}

func subscriber(channel <-chan *entity.ZipCode) (*entity.ZipCode, error) {
	select {
	case zipCodeInfo := <-channel:
		return zipCodeInfo, nil
	case <-time.After(time.Second):
		return nil, errors.New("time limit exceeded")
	}
}
