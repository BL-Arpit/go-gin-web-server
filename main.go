// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"runtime"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	ConfigRuntime()
// 	StartWorkers()
// 	StartGin()
// }

// // ConfigRuntime sets the number of operating system threads.
// func ConfigRuntime() {
// 	nuCPU := runtime.NumCPU()
// 	runtime.GOMAXPROCS(nuCPU)
// 	fmt.Printf("Running with %d CPUs\n", nuCPU)
// }

// // StartWorkers start starsWorker by goroutine.
// func StartWorkers() {
// 	go statsWorker()
// }

// // StartGin starts gin web server with setting router.
// func StartGin() {
// 	gin.SetMode(gin.ReleaseMode)

// 	router := gin.New()
// 	router.Use(rateLimit, gin.Recovery())
// 	router.LoadHTMLGlob("resources/*.templ.html")
// 	router.Static("/static", "resources/static")
// 	router.GET("/", index)
// 	router.GET("/room/:roomid", roomGET)
// 	router.POST("/room-post/:roomid", roomPOST)
// 	router.GET("/stream/:roomid", streamRoom)

//		port := os.Getenv("PORT")
//		if port == "" {
//			port = "8080"
//		}
//		if err := router.Run(":" + port); err != nil {
//	        log.Panicf("error: %s", err)
//		}
//	}
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
