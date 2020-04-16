// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package midware

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

func SwggerWrapper(f func(engine *gin.Engine)) func (engine *gin.Engine) {
    return func(engine *gin.Engine) {
        engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }
}
