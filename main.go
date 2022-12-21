package main

import (
	"resource_be/application/routes"

	"github.com/gin-gonic/gin"
)
var router = gin.Default()

func main() { 
	routes.Run( )
}