package engine

import "github.com/gin-gonic/gin"

type Engine struct {
	*gin.Engine
}

func NewEngine() *Engine {
	return &Engine{gin.Default()}
}
