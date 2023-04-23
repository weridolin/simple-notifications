package tools

import "fmt"

//目前可供选择的计划，暂时不支持自定义

type Period struct {
	Cron string
}

var (
	Minutely = "@minutely"
	Hourly   = "@hourly"
	Daily    = "@daily"
	Weekly   = "@weekly"
	Monthly  = "@monthly"
	Yearly   = "@yearly"
	HalfHour = "@every 30m"
	HalfDay  = "@every 12h"
)

func EveryDayAt(hour, minute int) string {
	return fmt.Sprintf("@every %dh%dm", hour, minute)
}

func EveryWeekAt(weekday, hour, minute int) string {
	return fmt.Sprintf("@every %dw%dh%dm", weekday, hour, minute)
}

func EveryMonthAt(day, hour, minute int) string {
	return fmt.Sprintf("@every %dm%dh%dm", day, hour, minute)
}
