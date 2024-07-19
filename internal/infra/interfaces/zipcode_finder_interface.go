package interfaces

import (
	"time"

	"github.com/brenddonanjos/multithreading_api/internal/entity"
)

type ZipcodeFinderInterface interface {
	FetchZipCode(zipCode string, startTime time.Time) (*entity.ZipCode, error)
}
