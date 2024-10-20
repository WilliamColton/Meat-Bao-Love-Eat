package middleware

import "github.com/gin-gonic/gin"

func JWTAuthMiddleware(j *JWT) gin.HandlerFunc {
	return func(context *gin.Context) { //返回的不是函数，而是函数定义（因为匿名函数末尾没加括号）
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, "need token")
			context.Abort()
			return
		}

		claims, err := j.ValidateJWT(tokenString)
		if err != nil {
			context.JSON(401, "token invalid")
			context.Abort()
			return
		}

		context.Set("claims", claims) //把claims存入claims键值对中，后续只需要对其进行类型断言（断言为Claims）即可
		context.Next()
	}
}
