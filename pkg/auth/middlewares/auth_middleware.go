package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leonardonatali/file-metadata-api/pkg/auth"
	"github.com/leonardonatali/file-metadata-api/pkg/auth/dto"
	"github.com/leonardonatali/file-metadata-api/pkg/users/entities"
)

var AuthMiddleware = func(c *gin.Context) {
	var dto dto.AuthDto

	if err := c.ShouldBindHeader(&dto); err != nil {
		log.Printf("cannot bind authorization header: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "token must be provided"})
		return
	}

	dto.Token = strings.TrimSpace(dto.Token)

	if dto.Token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token must be provided"})
		return
	}

	//Adiciona o novo usuário ao contexto da request
	ctx := context.WithValue(c.Request.Context(), auth.ContextUserKey, entities.User{
		Token: dto.Token,
	})

	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
