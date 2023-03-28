package web

import (
	"fmt"
	//"io/ioutil"
  "log"
	// "time"
  "encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  "github.com/imroc/req/v3"
)


func Server() {
  e := echo.New()

 // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  if err := e.Start(":8080"); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}

type UUIDResponse struct {
  UUID string
}

func (r *UUIDResponse) String() string {
  return r.UUID
}

// Handler
func hello(c echo.Context) error {
  client := req.C() // Use C() to create a client.
	resp, err := client.R().Get("https://httpbin.org/uuid")
	if err != nil {
		log.Fatal(err)
	}
  log.Println(resp.String())
  uuid := UUIDResponse{}
  fmt.Println(uuid)
	json.Unmarshal([]byte(resp.String()), &uuid)
  fmt.Println(uuid.String())
  return c.String(http.StatusOK, "Hello, World!")
}
