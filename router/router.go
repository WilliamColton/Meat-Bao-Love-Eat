package router

import (
	"Score_System/controller"
	"Score_System/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	controller *controller.Controller
	*gin.Engine
}

func NewRouter(controller *controller.Controller) *Router {
	router := gin.Default()

	router.Use(cors.Default())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/users")
			{
				user.POST("/login", controller.UserLogin)
				user.POST("/register", controller.UserRegister)
			}

			evaluation := v1.Group("/evaluations")
			{
				evaluation.GET("", controller.GetEvaluations)
			}

			level := v1.Group("/levels")
			{
				level.GET("", controller.GetLevels)
				level.GET("/id", controller.GetLevelByID)
				level.GET("/id/items", controller.GetItemsByLevelID)
			}

			item := v1.Group("/items")
			{
				item.GET("/id", controller.GetItemByID)
				item.GET("/id/evaluations", controller.GetEvaluationsByItemID)
				item.GET("/id/score", controller.GetScoreByItemID)
				item.GET("/random", controller.GetRandomItemsByLimit)
			}

			auth := v1.Group("/auth")
			auth.Use(middleware.JWTAuthMiddleware(controller.Jwt)) //记住此处写法，非常优美(我指jwt函数的返回值)
			{
				userWithAuth := auth.Group("/users")
				{
					userWithAuth.GET("/id/evaluations", controller.GetEvaluationsByUserID)
				}

				itemWithAuth := auth.Group("/items")
				{
					itemWithAuth.POST("", controller.CreateItem)
				}

				evaluationWithAuth := auth.Group("/evaluations")
				{
					evaluationWithAuth.GET("/id", controller.GetEvaluationByID)

					evaluationWithAuth.POST("", controller.CreateEvaluation)
				}

				levelWithAuth := auth.Group("/levels")
				{
					levelWithAuth.POST("", controller.CreateLevel)
				}
			}
		}
	}

	return &Router{
		controller: controller,
		Engine:     router,
	}
}
