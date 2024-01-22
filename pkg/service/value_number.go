package service

import (
	"context"
	"log"
	"strconv"
	"trafilea/cmd/logger"
)

type ValueNumberService struct{}

func NewValueNumberService() ValueNumberService {
	return ValueNumberService{}
}

func (s *ValueNumberService) Process(ctx context.Context, number int) string {
	logger.Info(ctx, "starting...")

	return ProcessValue(number)
}

func ProcessValue(i int) string {
	numberType := divisibleMap[i%15]

	if numberType == "" {
		numberType = strconv.Itoa(i)
	}

	log.Println(numberType)

	return numberType
}
