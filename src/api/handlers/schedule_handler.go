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
	// g.PUT("/:id", updateSchedule)
	// g.DELETE("/:id", deleteSchedule)
}

func createSchedule(c echo.Context)(err error){
	srv := api.GetAPICalendar()
	event := &calendar.Event{
	Summary: "Google I/O 2015",
	Location: "800 Howard St., San Francisco, CA 94103",
	Description: "A chance to hear more about Google's developer products.",
	Start: &calendar.EventDateTime{
		DateTime: "2021-06-25T10:15:00+09:00",
		TimeZone: "America/Los_Angeles",
	},
	End: &calendar.EventDateTime{
		DateTime: "2021-06-25T012:15:00+09:00",
		TimeZone: "America/Los_Angeles",
	},
	Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
	Attendees: []*calendar.EventAttendee{
	},
}
	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
	log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
	return c.JSON(http.StatusOK, event)
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
        fmt.Println("Upcoming events:")
        if len(events.Items) == 0 {
                fmt.Println("No upcoming events found.")
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

// func updateSchedule(c echo.Context) (err error) {
// 	db := db.InitDB()
// 	newData := models.User{}
// 	err = c.Bind(&newData)
// 	originData := &models.User{}
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	db.Where("id = ?", id).First(originData)
// 	fmt.Println("このログ出力したい", newData.ID)
// 	if newData.ID != 0 && newData.ID != originData.ID {
// 		return c.JSON(http.StatusConflict, "urlに含まれるIDと編集データIDが一致していません")
// 	}
// 	if (newData.CreatedAt != time.Time{}) && (newData.CreatedAt != originData.CreatedAt) {
// 		return c.JSON(http.StatusConflict, "作成時間の編集はできません")
// 	}
// 	if (newData.UpdatedAt != time.Time{}) && (newData.UpdatedAt != originData.UpdatedAt) {
// 		return c.JSON(http.StatusConflict, "更新時間の編集はできません")
// 	}
// 	if newData.Name != originData.Name {
// 		originData.Name = newData.Name
// 	}
// 	if newData.Email != originData.Email {
// 		originData.Email = newData.Email
// 	}
// 	db.Save(&originData)

// 	return c.JSON(http.StatusOK, newData)
// }

// func deleteSchedule(c echo.Context) (err error) {
// 	db := db.InitDB()
// 	data := &models.User{}
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	db.Where("id = ?", id).First(data)
// 	db.Delete(&data)
// 	return c.JSON(http.StatusOK,data)
// }
