package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld2(t *testing.T) {

	w := httptest.NewRecorder()
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	router.GET("/hello", HelloWorld)

	t.Run("get data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}
