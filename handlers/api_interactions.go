package handlers

import (
	"github.com/gin-gonic/gin"
)

func LigthsOff(c *gin.Context) {
	masterId := c.Param("masterId")
	mqtt_client.Publish(masterId, 0, false, "ligths_off")
}

func LigthsOn(c *gin.Context) {
	masterId := c.Param("masterId")
	mqtt_client.Publish(masterId, 0, false, "ligths_on")
}

func StaticOff(c *gin.Context) {
	masterId := c.Param("masterId")
	mqtt_client.Publish(masterId, 0, false, "static_off")
}

func StaticOn(c *gin.Context) {
	masterId := c.Param("masterId")
	mqtt_client.Publish(masterId, 0, false, "static_on")
}

func SetColor(c *gin.Context) {
	masterId := c.Param("masterId")
	mqtt_client.Publish(masterId, 0, false, "static_on")
}

func StartOption(c *gin.Context) {
	masterId := c.Param("masterId")
	mqtt_client.Publish(masterId, 0, false, "static_on")
}
