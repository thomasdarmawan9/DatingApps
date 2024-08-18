package routes

import (
	"final-project/controllers"
	_ "final-project/docs" // Import the generated docs
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Perbarui impor ini
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Dating Apps Documentation API
// @version 2.0
// @description
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User routes
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)

		// Routes after user creation requires authentication and restriction middleware
		userRouter.Use(middlewares.Authentication())
		userRouter.Use(middlewares.RestrictMiddleware())
		userRouter.PUT("/:id", controllers.UpdateUser)
		userRouter.DELETE("/:id", controllers.DeleteUser)
	}

	// Photo routes
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.Use(middlewares.RestrictMiddleware())

		photoRouter.GET("/", controllers.GetAllPhotos)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/:photoId", controllers.GetOnePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	// Social Media routes
	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.Use(middlewares.RestrictMiddleware())

		socialMediaRouter.GET("/", controllers.GetAllSocialMedias)
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", controllers.GetOneSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	// Match Profile routes (new)
	matchProfileRouter := r.Group("/swipe")
	{
		matchProfileRouter.Use(middlewares.Authentication())
		matchProfileRouter.Use(middlewares.RestrictMiddleware())

		matchProfileRouter.POST("/:profileID/:otherProfileID", controllers.SwipeProfile)
	}

	return r
}
