package main

import (
	"Resource.Application/routes"

	"github.com/gin-gonic/gin"
)
var router = gin.Default()

func main() { 
	routes.Run( )
}