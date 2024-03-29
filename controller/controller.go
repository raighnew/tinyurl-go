package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raighneweng/tinyurl-go/pkg/app"
	"github.com/raighneweng/tinyurl-go/pkg/e"
	"github.com/raighneweng/tinyurl-go/pkg/gredis"
	"github.com/raighneweng/tinyurl-go/pkg/murmurhash"
)

type TinyUrlBody struct {
	Url string `json:"url" form:"url" valid:"Required;MaxSize(100)"`
}

func Generate(c *gin.Context) {
	var (
		appG    = app.Gin{C: c}
		uParams TinyUrlBody
	)

	httpCode, errCode := app.BindAndValid(c, &uParams)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	seed, err := gredis.Incr("TINY_URL_SEED")

	if err != nil {
		log.Print(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	hashString := murmurhash.GenerateMurmurHash(uParams.Url, seed)

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"seed": hashString,
	})
}
