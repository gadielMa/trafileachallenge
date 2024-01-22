package service

import (
	"context"
	"trafilea/cmd/logger"
)

type DivisibleService struct{}

func NewDivisibleService() DivisibleService {
	return DivisibleService{}
}

var divisibleMap = map[int]string{3: "Type 1", 5: "Type 2", 6: "Type 1", 9: "Type 1", 10: "Type 2", 12: "Type 1", 0: "Type 3"}

func (s *DivisibleService) Prints(ctx context.Context) []string {
	logger.Info(ctx, "starting...")

	var listOfDivisible []string

	for i := 1; i <= 100; i++ {
		numberType := ProcessValue(i)
		listOfDivisible = append(listOfDivisible, numberType)
	}

	return listOfDivisible
}
