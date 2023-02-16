package cronx

// 每5秒
func Every5s() string {
	return "*/5 * * * * *"
}

// 每分钟
func EveryMinute() string {
	return "0 */1 * * * *"
}

// 每五分钟
func EveryFiveMinute() string {
	return "0 */5 * * * *"
}

// 每十分钟
func EveryTenMinute() string {
	return "0 */10 * * * *"
}

// 每半小时
func EveryHalfHour() string {
	return "0 0,30 * * * *"
}

// 每几分钟执行 1,2,3,26
func Hourly(m string) string {
	return "0 " + m + " * * * *"
}

// 每天几点执行
func Daily(h string) string {
	return "0 0 " + h + " * * *"
}
