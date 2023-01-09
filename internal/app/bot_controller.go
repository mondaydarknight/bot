package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/molpadia/molpastream/configs"
	"github.com/molpadia/molpastream/internal/domain/repository"
	"github.com/molpadia/molpastream/internal/infrastructure"
	"github.com/molpadia/molpastream/internal/service"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BotController struct {
	repo repository.BotRepository
	serv *service.BotService
}

// Create a new controller instance.
func NewBotController(ctx context.Context) *BotController {
	repo := infrastructure.NewBotRepository(ctx, options.Client().ApplyURI(configs.AppConfig.MongoUri))
	return &BotController{repo, service.NewBotService(repo)}
}

// Get a list of push messages.
func (c *BotController) get(ctx *gin.Context) {
	var messages []Message
	msgs, err := c.repo.Get(ctx.Query("userId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	for _, msg := range msgs {
		messages = append(messages, Message{msg.UserId, msg.Text, msg.CreatedAt})
	}
	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
}

// Send a push message to LINE bot.
func (c *BotController) send(ctx *gin.Context) {
	var msg Message
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	if err := c.serv.Send(msg.UserId, msg.Text); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

// Store a push message to the persistence via webhook.
func (c *BotController) webhook(ctx *gin.Context) {
	bot, _ := linebot.New(configs.AppConfig.LineChannelSecret, configs.AppConfig.LineAccessToken)
	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	if err = c.serv.Webhook(events); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{})
}
