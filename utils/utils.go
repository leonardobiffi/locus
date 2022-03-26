package utils

import (
	"fmt"
	"time"
)

func FormatResponseTime(start time.Time) string {
	return fmt.Sprintf("%.3f ms", time.Since(start).Seconds())
}
