package tools

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

//目前可供选择的计划，暂时不支持自定义

type Period struct {
	Cron string
}

var (
	Minutely = "* * */1 * * ?"
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

func CronValidator(f validator.FieldLevel) bool {
	value := f.Field().String()
	pattern := `(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(value)

}
