package tableprint

import "time"

type Timezone string

const (
	Timezone_Local = "local"
	Timezone_UTC   = "utc"
)

func formatTime(t time.Time, timezone Timezone) string {
	switch timezone {
	case Timezone_Local:
		return t.Local().String()
	case Timezone_UTC:
		return t.UTC().String()
	default:
		panic("unknown timezone: " + timezone)
	}
}
