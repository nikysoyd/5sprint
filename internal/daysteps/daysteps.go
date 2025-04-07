package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("wrong data amount%d", len(parts))
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("steps error parsing: %w", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("duration error parsing: %v", err)
	}
	ds.Duration = duration

	return nil

}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration %v", ds.Duration)
	}

	if ds.Steps <= 0 {
		return "", fmt.Errorf("amount steps %d", ds.Steps)
	}

	distance := spentenergy.Distance(ds.Steps)

	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	message := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n", ds.Steps, distance, calories)

	return message, nil

}
