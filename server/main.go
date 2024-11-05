package main

import (
	"forum/config"
	"forum/database"
	"forum/handler"
	"forum/model"
	"forum/store"
	"forum/utility"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config, err := config.Load("./config.json")
	if err != nil {
		log.Fatal("failed to load config file: ", err)
	}
	db, err := database.Init(config.DSN)
	if err != nil {
		log.Fatal("failed to initialize database: ", err)
	}
	err = db.AutoMigrate(
		&model.Account{},
		&model.Thread{},
		&model.Subforum{},
		&model.Reply{},
	)
	if err != nil {
		log.Fatal("failed to initialize database: ", err)
	}

	server := echo.New()
	server.Validator = &utility.Validator{Validator: validator.New()}
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())

	cookieStore := sessions.NewCookieStore()
	cookieStore.Options = &sessions.Options{
		MaxAge:   24 * 60 * 60 * 1000,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	apiRouter := server.Group("/api")

	accountStore := store.NewAccountStore(db)
	accountHandler := handler.NewAccountHandler(accountStore)
	AddAccountRoutes(apiRouter, accountHandler)

	subforumStore := store.NewSubforumStore(db)
	subforumHandler := handler.NewSubforumHandler(subforumStore)
	AddSubforumRoutes(apiRouter, subforumHandler)

	threadStore := store.NewThreadStore(db)
	threadHandler := handler.NewThreadHandler(threadStore)
	AddThreadRoutes(apiRouter, threadHandler)

	replyStore := store.NewReplyStore(db)
	replyHandler := handler.NewReplyHandler(replyStore)
	AddReplyRoutes(apiRouter, replyHandler)

	authHandler := handler.NewAuthHandler(accountStore, cookieStore)
	AddAuthRoutes(apiRouter, authHandler)

	err = server.Start(config.Host + ":" + config.Port)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}

func AddAccountRoutes(
	router *echo.Group,
	accountHandler *handler.AccountHandler,
) {
	accountRouter := router.Group("/account")
	accountRouter.GET("/", accountHandler.GetAccounts)
	accountRouter.GET("/:accountId", accountHandler.GetAccount)
	accountRouter.POST("/", accountHandler.PostAccount)
	accountRouter.DELETE("/:accountId", accountHandler.DeleteAccount)
}

func AddSubforumRoutes(
	router *echo.Group,
	subforumHandler *handler.SubforumHandler,
) {
	subforumRouter := router.Group("/subforum")
	subforumRouter.GET("/", subforumHandler.GetSubforums)
	subforumRouter.GET("/:subforumId", subforumHandler.GetSubforums)
	subforumRouter.POST("/", subforumHandler.PostSubforum)
	subforumRouter.PATCH("/:subforumId", subforumHandler.PatchSubforum)
	subforumRouter.DELETE("/:subforumId", subforumHandler.DeleteSubforum)
}

func AddThreadRoutes(
	router *echo.Group,
	threadHandler *handler.ThreadHandler,
) {
	threadRouter := router.Group("/thread")
	threadRouter.GET("/", threadHandler.GetThreads)
	threadRouter.GET("/:threadId", threadHandler.GetThread)
	threadRouter.POST("/", threadHandler.PostThread)
	threadRouter.PATCH("/:threadId", threadHandler.PatchThread)
	threadRouter.DELETE("/:threadId", threadHandler.DeleteThread)
}

func AddReplyRoutes(
	router *echo.Group,
	replyHandler *handler.ReplyHandler,
) {
	replyRouter := router.Group("/reply")
	replyRouter.GET("/", replyHandler.GetReplies)
	replyRouter.GET("/:replyId", replyHandler.GetReply)
	replyRouter.POST("/", replyHandler.PostReply)
	replyRouter.PATCH("/:replyId", replyHandler.PatchReply)
	replyRouter.DELETE("/:replyId", replyHandler.DeleteReply)
}

func AddAuthRoutes(
	router *echo.Group,
	authHandler *handler.AuthHandler,
) {
	authRouter := router.Group("/auth")
	authRouter.POST("/login", authHandler.Login)
	authRouter.POST("/logout", authHandler.Logout)
}
