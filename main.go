package main

import (
  "github.com/gin-gonic/gin"
  "github.com/S9S99/todo_list/controllers"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*.tmpl")
  router.GET("/", controllers.GetIndex)

  router.Run()
}

