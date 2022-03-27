package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func main() {
//     // controller
//     http.HandleFunc("/", echoHello)
//     // port
//     http.ListenAndServe(":8000", nil)
// }

// func echoHello(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "<h1>Hello World</h1>")
// }

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hogefugapiyo")
	})

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	router.Run(":8000")
}
