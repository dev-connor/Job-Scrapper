package pkg

import (
	"strings"
	"time"
)

func TimeToKo(time time.Duration) string {
	timeStr := time.String()

	if strings.HasSuffix(timeStr, "s") {
		return strings.ReplaceAll(timeStr, "s", " 초")

	} else if strings.HasSuffix(timeStr, "m") {
		return strings.ReplaceAll(timeStr, "m", " 밀리 초")

	} else {
		return time.String()
	}
}
