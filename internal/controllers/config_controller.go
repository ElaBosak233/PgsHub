package controllers

import "github.com/gin-gonic/gin"

type ConfigController interface {
	Find(ctx *gin.Context)
}