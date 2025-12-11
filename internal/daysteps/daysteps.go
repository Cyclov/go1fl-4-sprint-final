package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"log"

	spentcalories "github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	separatorString := ","
	dataStrings := strings.Split(data, separatorString)
	dataDurtionPosition := 1
	dataStepsPosition := 0

	if len(dataStrings) != 2 {
		return 0, time.Duration(0), errors.New("data exeption: wrong data string")
	}

	steps, err := strconv.Atoi(dataStrings[dataStepsPosition])

	if err != nil {
		return 0, time.Duration(0), err
	}

	if steps <= 0 {
		return 0, time.Duration(0), errors.New("data exeption: wrong number of steps")
	}

	duration, err := time.ParseDuration(dataStrings[dataDurtionPosition])

	if err != nil {
		return 0, time.Duration(0), err
	}

	if duration <= time.Duration(0) {
		return 0, time.Duration(0), errors.New("data exeption: wrong number of duration")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)

	if err != nil {
		log.Println(err)

		return ""
	}

	distance := float64(steps) * stepLength / mInKm
	kkal, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		log.Println(err)

		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, kkal)

	return result

}
