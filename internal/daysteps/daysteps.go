package daysteps

import (
	"errors"
	"fmt"
	"internal/spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65gi
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	separatorString := ","
	dataStrings := strings.Split(data, separatorString)
	dataDurtionPosition := 0
	dataStepsPosition := 1

	if len(dataStrings) < 2 {
		return 0, time.Duration(0), errors.New("data exeption: wrong data string")
	}

	steps, err := strconv.Atoi(dataStrings[dataStepsPosition])

	if err != nil {
		return 0, time.Duration(0), err
	}

	if steps < 0 {
		return 0, time.Duration(0), errors.New("data exeption: wrong number of steps")
	}

	duration, err := time.ParseDuration(dataStrings[dataDurtionPosition])

	if err != nil {
		return 0, time.Duration(0), err
	}

	if duration == time.Duration(0) {
		return 0, time.Duration(0), errors.New("data exeption: wrong number of duration")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage()

	if err != nil {
		fmt.Println(err)

		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	kkal := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %f км.\nВы сожгли %f ккал.\n", steps, distance, kkal)

	return result

}
