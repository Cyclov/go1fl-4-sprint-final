package spentcalories

import (
	"time"
	"fmt"
	"strings"
	"errors"
	"strconv"
	"log"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	separatorString := ","
	dataStrings := strings.Split(data, separatorString)
	DurtionPosition := 2
	TrainingPosition := 1
	StepsPosition := 0

	if len(dataStrings) != 3 {
		return 0,"", time.Duration(0), errors.New("data exeption: wrong data string")
	}

	steps, err := strconv.Atoi(dataStrings[StepsPosition])

	if err != nil {
		return 0, "", time.Duration(0), err
	}

	if steps <= 0 {
		return 0, "", time.Duration(0), errors.New("data exeption: wrong number of steps")
	}

	duration, err := time.ParseDuration(dataStrings[DurtionPosition])

	if err != nil {
		return 0, "", time.Duration(0), err
	}

	if duration <= 0 {
		return 0, "", time.Duration(0), errors.New("data exeption: wrong number of duration")
	}

	trainingType := dataStrings[TrainingPosition]

	return steps, trainingType, duration, nil
}

func distance(steps int, height float64) float64 {
	return height*stepLengthCoefficient*float64(steps)/mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	
	if duration <= time.Duration(0) {return 0}

	return distance(steps,height)/ duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	var kal float64
	var err error 
	steps, trainingType, duration, err := parseTraining(data)
		
	if err != nil {
		log.Println(err)
		return "", err
	}

	distance := distance(steps, height)
	speed :=  meanSpeed(steps, height, duration)
	
	switch strings.ToUpper(trainingType) {
	case "БЕГ":  kal, err = RunningSpentCalories(steps, weight, height, duration)
	case "ХОДЬБА": kal, err = WalkingSpentCalories(steps, weight, height, duration)
	default: return "", errors.New("неизвестный тип тренировки") 	
	}
	
	if err != nil {
		log.Println(err)
		return "", err
	}	



	str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration.Hours(), distance, speed, kal)
	return str, nil 
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	
	if weight  <= 0 {return 0, errors.New("data exeption: wrong weight")} 
	if height  <= 0 {return 0, errors.New("data exeption: wrong height")} 
	if steps  <= 0 {return 0, errors.New("data exeption: wrong steps")} 
	if duration <= time.Duration(0) {return 0, errors.New("data exeption: wrong duration")} 

	return (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH , nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	kkal, err := RunningSpentCalories(steps, weight, height, duration)
	return kkal * walkingCaloriesCoefficient, err 


}
