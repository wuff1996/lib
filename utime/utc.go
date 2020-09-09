package utime

import (
	"fmt"
	"time"
)

func GetTIme()string{

	myTime := time.Now().UTC()
	return fmt.Sprintf(
		"%d%02d%02dT%02d%02d%02dZ",
		myTime.Year(),
		myTime.Month(),
		myTime.Day(),
		myTime.Hour(),
		myTime.Minute(),
		myTime.Second())
}
