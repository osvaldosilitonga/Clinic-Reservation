package helpers

import "time"

func ConvertTimeLocal(t int64) time.Time {
	return time.Unix(t, 0).Local()
}
