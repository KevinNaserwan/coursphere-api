package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	postgresConfig "github.com/kevinnaserwan/coursphere-api/config/postgres"
	bookController "github.com/kevinnaserwan/coursphere-api/internal/controller/book"
	bookCategoryController "github.com/kevinnaserwan/coursphere-api/internal/controller/book_category"
	userController "github.com/kevinnaserwan/coursphere-api/internal/controller/user"
	http "github.com/kevinnaserwan/coursphere-api/internal/http/server"
	"github.com/kevinnaserwan/coursphere-api/internal/http/server/middleware"
	BookRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/book/postgres"
	BookCategoryRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/book_category/postgres"
	UserRepository "github.com/kevinnaserwan/coursphere-api/internal/repository/user/postgres"
	BookService "github.com/kevinnaserwan/coursphere-api/internal/service/book"
	BookCategoryService "github.com/kevinnaserwan/coursphere-api/internal/service/book_category"
	UserService "github.com/kevinnaserwan/coursphere-api/internal/service/user"
	"github.com/kevinnaserwan/coursphere-api/internal/util/jwt"
	"github.com/kevinnaserwan/coursphere-api/internal/util/postgres"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func (a *app) StartServer() {
	server := http.NewHTTPServer(a.config.Environment)

	server.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"message": "Welcome to coursphere API",
		})
	})

	var db *gorm.DB
	if a.config.Environment == "release" {
		db = postgres.NewDB(a.config.DbReleaseURL)
	} else if a.config.Environment == "test" {
		db = postgres.NewDB(a.config.DbTestURL)
	} else {
		db = postgres.NewDB(a.config.DbDebugURL)
	}

	postgresConfig.Migrate(db)

	if a.config.Environment != "release" {
		docs.SwaggerInfo.Title = "COURSPHERE API"
		docs.SwaggerInfo.Description = "E-Commerce Mobile API"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", a.config.Port)
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}

		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	jwtManager := jwt.NewJWTManager(a.config.JwtSecret)
	mailDialer := gomail.NewDialer(
		a.config.EmailHost,
		int(a.config.EmailPort),
		a.config.EmailUser,
		a.config.EmailPassword,
	)

	initController(server, db, jwtManager, mailDialer)

	log.Printf("Server is running on %s mode", strings.ToUpper(a.config.Environment))
	err := server.Run(fmt.Sprintf(":%d", a.config.Port))
	if err != nil {
		panic(err)
	}

}

func initController(
	root *gin.Engine,
	db *gorm.DB,
	jwtManager *jwt.JWTManager,
	mailDialer *gomail.Dialer,
) {
	userRepository := UserRepository.NewUserRepository(db)
	bookCategoryRepository := BookCategoryRepository.NewBookCategoryRepository(db)
	bookRepository := BookRepository.NewBookRepository(db)
	userService := UserService.NewUserService(userRepository, jwtManager, mailDialer)
	bookCategoryService := BookCategoryService.NewBookCategoryService(bookCategoryRepository)
	bookService := BookService.NewBookService(bookRepository)

	routerGroup := root.Group("/api/v1")

	routerGroup.Use(middleware.ErrorHandler())

	userController.NewUserController(routerGroup.Group("/user"), userService, jwtManager)
	bookCategoryController.NewBookCategoryController(routerGroup.Group("/book-category"), bookCategoryService)
	bookController.NewBookController(routerGroup.Group("/book"), bookService)
}
