package main

import (
	"LWS_Web/models"
	_ "LWS_Web/routers"
	"LWS_Web/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/patrickmn/go-cache"
)

func main() {
	models.Init()
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	beego.Run()
}
