package utils

import "strconv"

func Pad(n, size int) string {
	str := strconv.Itoa(n)
	for len(str) < size {
		str = "0" + str
	}
	return str
}

func FormatTime(rawSeconds int) string {
	hours := rawSeconds / 3600
	minutes := (rawSeconds % 3600) / 60
	seconds := (rawSeconds % 3600) % 60
	if hours == 0 {
		return Fmt("%s:%s", Pad(minutes, 2), Pad(seconds, 2))
	}
	return Fmt("%s:%s:%s",
		Pad(hours, 2), Pad(minutes, 2), Pad(seconds, 2),
	)
}
