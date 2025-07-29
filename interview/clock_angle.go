package interview

import (
	"fmt"
	"math"
)

func clockAngle(hour, minute int) float64 {
	hourAngle := 0.5 * float64(60*hour+minute)
	minuteAngle := 6.0 * float64(minute)
	angle := math.Abs(hourAngle - minuteAngle)
	if angle > 180 {
		angle = 360 - angle
	}
	return angle
}

func ClockAngle() {
	hour := 3
	minute := 30
	fmt.Printf("Angle: %.2f degrees\n", clockAngle(hour, minute))
	fmt.Println("")
}
