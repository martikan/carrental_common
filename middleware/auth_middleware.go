package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/martikan/carrental_common/exception"
	"github.com/martikan/carrental_common/util"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker util.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is empty")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.ApiError(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.ApiError(err))
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type: %v", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.ApiError(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.ApiError(err))
			return
		}

		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
