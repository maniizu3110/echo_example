package handlers

import (
	"fmt"
	"log"
	"myapp/src/api/api"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/calendar/v3"
)

func ScheduleHandler(g *echo.Group) {
	g.POST("/:id", createSchedule)
	g.GET("/:id", getSchedule)
	g.PUT("/:id", updateSchedule)
	// g.DELETE("/:id", deleteSchedule)
}

func createSchedule(c echo.Context)(err error){
	//TODO:タイムゾーンなどデフォルトを設定
	//TODO:フロントから受け取ったものをここで当てはめていくことで使えるデータをgoogle apiに挿入
	srv := api.GetAPICalendar()
	event := &calendar.Event{
	Summary: "Google I/O 2015",
	Location: "800 Howard St., San Francisco, CA 94103",
	Description: "A chance to hear more about Google's developer products.",
	Start: &calendar.EventDateTime{
		DateTime: "2015-05-28T09:00:00-07:00",
		TimeZone: "America/Los_Angeles",
	},
	End: &calendar.EventDateTime{
		DateTime: "2015-05-28T17:00:00-07:00",
		TimeZone: "America/Los_Angeles",
	},
	Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
	Attendees: []*calendar.EventAttendee{
		&calendar.EventAttendee{Email:"lpage@example.com"},
		&calendar.EventAttendee{Email:"sbrin@example.com"},
	},
	}

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
	log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
	return c.JSON(http.StatusOK,event)
}
			
//UserHandler は日付を指定してやればその日付のstartとendを自動的に計算してその日の予定を返す
func getSchedule(c echo.Context) (err error) {
	//実際にはユーザーをIDで判別してそれに対応する予定を返す必要がある
	// id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		srv := api.GetAPICalendar()
		start := time.Now().Format(time.RFC3339)
        end := time.Now().AddDate(0,0,6).Format(time.RFC3339)
        events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(start).TimeMax(end).MaxResults(10).OrderBy("startTime").Do()
        if err != nil {
			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
        }
        if len(events.Items) == 0 {
			} else {
				for _, item := range events.Items {
					date := item.Start.DateTime
					if date == "" {
						date = item.Start.Date
					}
					fmt.Printf("%+v \n", item.Summary)
                }
			}
			data := events.Items
			
			
			return c.JSON(http.StatusOK, data)
		}
		
func updateSchedule(c echo.Context) (err error) {
	srv := api.GetAPICalendar()
	//jsonでフロントから構造をそのまま受け取ってeventにbindする
	event := new(calendar.Event)
	if err = c.Bind(&event); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//TODO:フロントから指定可能にする
	calendarId := "primary"
	event, err = srv.Events.Update(calendarId,event.Id,event).Do()
	return c.JSON(http.StatusOK, event)
}

// func deleteSchedule(c echo.Context) (err error) {
// 	db := db.InitDB()
// 	data := &models.User{}
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	db.Where("id = ?", id).First(data)
// 	db.Delete(&data)
// 	return c.JSON(http.StatusOK,data)
// }
