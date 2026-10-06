package classParse

import (
	"homewormanager/internal/errs"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/teambition/rrule-go"

	"homewormanager/internal/entity"
	"homewormanager/internal/utils"
)

type Class struct {
	Summary     string
	Start       time.Time
	End         time.Time
	Description string
	Location    string
	Category    string
	Dates       []time.Time
}

func IcalParse(icalLink string) ([]entity.Class, []string, error) {

	resp, err := http.Get(icalLink)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	bodyString := string(body)

	if len(bodyString) == 0 {
		return nil, nil, errs.NoIcal
	}

	bodyString = strings.Replace(bodyString, "X-SCHEDULE_VERSION-ID:", "X-SCHEDULE-VERSION-ID:", -1)

	calendar, err := ics.ParseCalendar(strings.NewReader(bodyString))
	if err != nil {
		return nil, nil, err
	}

	var classes []Class
	for _, event := range calendar.Events() {

		if utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertyCategories)).Value == "" {
			continue
		}

		class := Class{
			Summary:     utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertySummary)).Value,
			Description: utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertyDescription)).Value,
			Location:    utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertyLocation)).Value,
			Category:    utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertyCategories)).Value,
		}
		class.Start, err = time.ParseInLocation("20060102T150405", utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertyDtStart)).Value, time.Local)

		class.End, err = time.ParseInLocation("20060102T150405", utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertyDtEnd)).Value, time.Local)

		RRULE := utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertyRrule)).Value

		if RRULE != "" {
			rule, err := rrule.StrToRRule(RRULE)
			rule.DTStart(class.Start)
			if err != nil {
				return nil, nil, err
			}
			prepareDates := rule.All()

			exDateStr := utils.DeRef[ics.IANAProperty](event.GetProperty(ics.ComponentPropertyExdate)).Value

			exMap := make(map[time.Time]struct{})
			exDatesStr := strings.Split(exDateStr, ",")
			for _, date := range exDatesStr {
				if date == "" {
					continue
				}
				date, err := time.ParseInLocation("20060102T150405", date, time.Local)
				if err != nil {
					return nil, nil, err
				}
				exMap[date] = struct{}{}
			}
			for _, date := range prepareDates {
				if _, ok := exMap[date]; ok {
					continue
				}
				class.Dates = append(class.Dates, date)
			}
		} else {
			class.Dates = append(class.Dates, class.Start)
		}

		classes = append(classes, class)
	}

	subjects := parseSubject(classes)

	var result []entity.Class

	for _, class := range classes {

		duration := class.End.Sub(class.Start)
		for _, date := range class.Dates {
			result = append(result, entity.Class{
				StartTime:   date,
				EndTime:     date.Add(duration),
				Summary:     class.Summary,
				Description: class.Description,
				Location:    class.Location,
				Category:    entity.CategoryToNumber[class.Category],
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		startI := result[i].StartTime
		startJ := result[j].EndTime
		return startI.Before(startJ)
	})

	return result, subjects, nil
}
