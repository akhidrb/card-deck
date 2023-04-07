package router

import (
	"github.com/akhidrb/toggl-cards/pkg/controllers"
	controllersI "github.com/akhidrb/toggl-cards/pkg/interfaces/controllers"
	"github.com/akhidrb/toggl-cards/pkg/repositories"
	"github.com/akhidrb/toggl-cards/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitDeckDependencies(apiV1 *gin.RouterGroup, dbConn *gorm.DB) {
	ctrl := initDeckController(dbConn)
	deckInitRoutes(apiV1, ctrl)
}

func initDeckController(dbConn *gorm.DB) controllersI.IDeck {
	repo := repositories.NewDeck(dbConn)
	service := services.NewDeck(repo)
	return controllers.NewDeck(service)
}

func deckInitRoutes(apiGroup *gin.RouterGroup, ctrl controllersI.IDeck) {
	deckGroup := apiGroup.Group("/deck")
	deckGroup.POST("", ctrl.Create)
	deckGroup.GET("/:id", ctrl.GetByID)
	deckGroup.PUT("/:id/draw", ctrl.DrawCards)
}
