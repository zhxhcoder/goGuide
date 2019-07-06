package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"os"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, golang!</h1>
</body>
</html>
`))

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := gin.Default()
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.Run(":8083")
}
