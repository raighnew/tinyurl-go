package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raighneweng/tinyurl-go/pkg/app"
	"github.com/raighneweng/tinyurl-go/pkg/e"
	"github.com/raighneweng/tinyurl-go/pkg/gredis"
	"github.com/raighneweng/tinyurl-go/pkg/murmurhash"
	"github.com/raighneweng/tinyurl-go/pkg/setting"
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
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	hashString := murmurhash.GenerateMurmurHash(uParams.Url, seed)

	hashExist := gredis.Exists(hashString)

	if hashExist {
		hashString = murmurhash.GenerateMurmurHash(uParams.Url, seed)
	}

	err = gredis.Set(hashString, uParams.Url, 2*24*3600)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"url": fmt.Sprintf("%s/%s", setting.ServerSetting.BaseUrl, hashString),
	})
}

func GetUrl(c *gin.Context) {
	appG := app.Gin{C: c}
	urlHash := c.Param("urlHash")

	var resInString string

	_, err := gredis.Get(urlHash, &resInString)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"url": resInString,
	})
}
