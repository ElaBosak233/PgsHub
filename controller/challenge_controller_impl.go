package controller

import (
	"github.com/elabosak233/pgshub/model/data"
	"github.com/elabosak233/pgshub/service"
	"github.com/elabosak233/pgshub/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChallengeControllerImpl struct {
	challengeService service.ChallengeService
}

func NewChallengeController(appService service.AppService) ChallengeController {
	return &ChallengeControllerImpl{
		challengeService: appService.ChallengeService,
	}
}

func (c *ChallengeControllerImpl) Create(ctx *gin.Context) {
	createChallengeRequest := data.Challenge{}
	err := ctx.ShouldBindJSON(&createChallengeRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  utils.GetValidMsg(err, &createChallengeRequest),
		})
		return
	}
	_ = c.challengeService.Create(createChallengeRequest)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (c *ChallengeControllerImpl) Update(ctx *gin.Context) {
	var updateChallengeRequest data.Challenge
	err := ctx.ShouldBindJSON(&updateChallengeRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  utils.GetValidMsg(err, &updateChallengeRequest),
		})
		return
	}
	err = c.challengeService.Update(updateChallengeRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "更新失败",
		})
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (c *ChallengeControllerImpl) Delete(ctx *gin.Context) {
	deleteChallengeRequest := ChallengeDeleteRequest{}
	err := ctx.ShouldBindJSON(&deleteChallengeRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  utils.GetValidMsg(err, &deleteChallengeRequest),
		})
		return
	}
	err = c.challengeService.Delete(deleteChallengeRequest.Id)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "删除失败",
		})
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (c *ChallengeControllerImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	challengeData := c.challengeService.FindById(id)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": challengeData,
	})
}

func (c *ChallengeControllerImpl) FindAll(ctx *gin.Context) {
	challengeData := c.challengeService.FindAll()
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": challengeData,
	})

}