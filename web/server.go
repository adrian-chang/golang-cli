package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "time"

	"github.com/gofiber/fiber/v2"
)

type foo interface {
  read() int
}

type foo2 interface {
  foo
}

func Server() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
      fmt.Println("start")
      res := make(chan string)
      for i := 0; i < 10; i++ {
          go func(i int, c chan string) {
            resp, err := http.Get("http://example.com/")
            defer resp.Body.Close()
            ioutil.ReadAll(resp.Body)
            //fmt.Printf("%v\n", i)
            //c <- string(responseRead) 
            c <- string(fmt.Sprintf("%v", i))
            if err != nil {
            }
          }(i, res)
      }
      for loop := true; loop; {
        select {
        case v := <- res:
            fmt.Println(v)
            if v == "9" {
              loop = false
            }
        }
      }
      return c.SendString("Hello, World 👋!")
    })

    app.Listen(":3000")

}
