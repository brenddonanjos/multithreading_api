package service

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/brenddonanjos/multithreading_api/internal/entity"
	"github.com/brenddonanjos/multithreading_api/internal/infra/interfaces"
)

func GetZipCodeInfo(zipCode string) (*entity.ZipCode, error) {
	//Usa o contexto para limitar o tempo de execução da requisição para a api externa
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	zipCopdeFinderInterfaces := []interfaces.ZipcodeFinderInterface{
		NewBrasilApiService(ctx),
		NewViaCepService(ctx),
	}
	channel := make(chan *entity.ZipCode)
	wg := sync.WaitGroup{}

	for key, zipCodeFinderInterface := range zipCopdeFinderInterfaces {
		wg.Add(1)
		startTime := time.Now()
		go publisher(zipCode, zipCodeFinderInterface, channel, &wg, startTime, key)
	}
	var zipCodeInfo entity.ZipCode

	go subscriber(channel, &zipCodeInfo)
	wg.Wait()
	close(channel)

	return &zipCodeInfo, nil
}

func publisher(zipCode string, zipCodeFinderInterface interfaces.ZipcodeFinderInterface, channel chan<- *entity.ZipCode, wg *sync.WaitGroup, startTime time.Time, key int) {
	defer wg.Done()
	zipCodeInfo, err := zipCodeFinderInterface.FetchZipCode(zipCode, startTime)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Publisher ("+strconv.Itoa(key+1)+"): ", zipCodeInfo.Service, " - ", zipCodeInfo.ExecutionTime)
	channel <- zipCodeInfo
}

func subscriber(channel <-chan *entity.ZipCode, zipCodeInfo *entity.ZipCode) {
	for info := range channel {
		if zipCodeInfo == nil || zipCodeInfo.ZipCode == "" {
			*zipCodeInfo = *info
		}
	}
}
