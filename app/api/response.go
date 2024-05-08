package api

import "github.com/gin-gonic/gin"

func ResponseError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

func ResponseCreated(ctx *gin.Context, data interface{}) {
	ctx.JSON(201, gin.H{
		"code": 201,
		"data": data,
	})
}

func ResponseNoContent(ctx *gin.Context) {
	ctx.JSON(204, gin.H{
		"code": 204,
	})
}

func ResponseBadRequest(ctx *gin.Context, msg string) {
	ctx.JSON(400, gin.H{
		"code": 400,
		"msg":  msg,
	})
}

func ResponseUnauthorized(ctx *gin.Context, msg string) {
	ctx.JSON(401, gin.H{
		"code": 401,
		"msg":  msg,
	})
}

func ResponseForbidden(ctx *gin.Context, msg string) {
	ctx.JSON(403, gin.H{
		"code": 403,
		"msg":  msg,
	})
}

func ResponseNotFound(ctx *gin.Context, msg string) {
	ctx.JSON(404, gin.H{
		"code": 404,
		"msg":  msg,
	})
}

func ResponseInternalServerError(ctx *gin.Context, msg string) {
	ctx.JSON(500, gin.H{
		"code": 500,
		"msg":  msg,
	})
}
