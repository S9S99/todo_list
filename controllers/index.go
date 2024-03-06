package controllers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func GetIndex(ctx *gin.Context) {
  h := gin.H{
    "title": "TODO List",
  }

  ctx.HTML(http.StatusOK, "index.tmpl", h)
}
