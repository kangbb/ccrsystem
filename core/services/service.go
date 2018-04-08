package services

import (
	"regexp"
	"strconv"
	"time"
)

/***********************************SOME SUBFUNCTION*******************************/
func lessonToTime(lessonTime string) time.Time {
	mappingRegular := map[string]int{
		"一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9, "十": 1,
	}
	var lesson int
	reg := regexp.MustCompile(`[0-9\-]+`)
	date := reg.FindAllString(lessonTime, -1)[0]
	reg = regexp.MustCompile(`[\p{Han}]+`)
	hanstr := []rune(reg.FindAllString(lessonTime, -1)[0])
	switch len(hanstr) {
	case 3:
		lesson = mappingRegular[string(hanstr[1])]
		break
	case 4:
		lesson = mappingRegular[string(hanstr[1])]*10 + mappingRegular[string(hanstr[2])]
		break
	}
	startTime, _ := time.Parse("2006-01-02 15:04:05", date+" "+"08:00:00")
	duraTime := strconv.Itoa((lesson-1)*(45+10)) + "m"
	addTime, _ := time.ParseDuration(duraTime)
	return startTime.Add(addTime)
}
func timeToLesson(trueTime time.Time) string {
	mappingRegular := []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	var lesson string
	date := trueTime.Format("2006-01-02")
	startTime, _ := time.Parse("2006-01-02 15:04:05", date+" "+"08:00:00")
	lessonInNum := int(trueTime.Sub(startTime).Minutes())/55 + 1
	if lessonInNum > 10 {
		lesson = mappingRegular[9] + mappingRegular[lessonInNum%10-1]
	} else if lessonInNum == 10 {
		lesson = mappingRegular[9]
	} else {
		lesson = mappingRegular[lessonInNum%10-1]
	}

	return date + " " + "第" + lesson + "节"

}
