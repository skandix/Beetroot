package api

import (
	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

type news struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Message string `json:"Message"`
	Author  string `json:"Author"`
	Date    string `json:"Date"`
}

// TODO: ha date og data i samme, men det for bli senere.
type event struct {
	ID        string `json:"Id"`
	Name      string `json:"Name"`
	Location  string `json:"Location"`
	Date      string `json:"Date"`
	StartTime string `json:"Start Time"`
	Duration  string `json:"Duration"`
}

type job struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"URL"`
}

var nyheter = []news{
	{ID: "1", Title: "BetaLAN Pog!", Message: "Da er det straks klart for en ny runde med BetaLAN! BetaLAN vil foregå på bluebox som vanlig, og finner i år sted den 25.10. Husk å meld dere på arrangementet på forhånd slik at dere bare kan...", Author: "Clip", Date: "28.08.2020"},
	{ID: "2", Title: "CGI", Message: "21.09.2020 Kommer CGI for å holde bedriftspresentasjon med beta! Her er det gode muligheter for å bli kjent med en vikitg aktør fra næringslivet! Det blir servert Wraps og brus! Påmelding via docs...", Author: "Clip", Date: "28.08.2020"},
}

var events = []event{
	{ID: "1", Name: "Betalan", Location: "Blyatbox", Date: "20.08.2020", StartTime: "18:00", Duration: "72H"},
}

var jobs = []job{
	{ID: "1", Name: "TSOC", Description: "Vi søker deltids ansatte til TSOC", Url: "https://www.telenor.no/bedrift/sikkerhet/tsoc/"},
}

func GetNews(c *gin.Context) {
	c.JSON(http.StatusOK, nyheter)
}

func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, events)
}

func GetJobs(c *gin.Context) {
	c.JSON(http.StatusOK, jobs)
}

func GetTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"time":time.Now(),
	})
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Online",
	})
}

func StartAPI(Port int) {
	router := gin.Default()
    router.GET("/", root)
	v1 := router.Group("/v1")
	{
		events := v1.Group("/events")
		{
			events.GET("/", GetEvents)
		}
		jobs := v1.Group("/jobs")
		{
			jobs.GET("/", GetJobs)
		}
		news := v1.Group("/news")
		{
			news.GET("/", GetNews)
		}
		time := v1.Group("/time")
		{
			time.GET("/", GetTime)
		}
	}
	router.Run(fmt.Sprintf(":%d", Port))
}
