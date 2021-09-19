package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leonardonatali/file-metadata-api/pkg/auth"
	"github.com/leonardonatali/file-metadata-api/pkg/auth/dto"
	"github.com/leonardonatali/file-metadata-api/pkg/users"
	usersDto "github.com/leonardonatali/file-metadata-api/pkg/users/dto"
)

func GetAuthMiddleware(usersService *users.UsersService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto dto.AuthDto

		if err := c.ShouldBindHeader(&dto); err != nil {
			log.Printf("cannot bind authorization header: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "token must be provided"})
			return
		}

		dto.Token = strings.TrimSpace(dto.Token)

		if dto.Token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token must be provided"})
			return
		}

		// locaiza o usuário pelo Token
		user, err := usersService.GetUser(&usersDto.GetUserDto{
			Token: dto.Token,
		})

		if err != nil {
			panic(err)
		}

		if user == nil {
			//Caso não encontre um usuário com o token informado, cria um novo
			user, err = usersService.CreateUser(&usersDto.CreateUserDto{
				Token: dto.Token,
			})

			if err != nil {
				panic(err)
			}
		}

		//Adiciona o novo usuário ao contexto da request
		ctx := context.WithValue(c.Request.Context(), auth.ContextUserKey, user)

		log.Printf("%+v", user)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
