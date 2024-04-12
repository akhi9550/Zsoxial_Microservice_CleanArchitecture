package di

import (
	server "github.com/akhi9550/api-gateway/pkg/api"
	"github.com/akhi9550/api-gateway/pkg/api/handler"
	"github.com/akhi9550/api-gateway/pkg/client"
	"github.com/akhi9550/api-gateway/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {
	authClient := client.NewAuthClient(cfg)
	authHandler := handler.NewAuthHandler(authClient)

	postClient := client.NewPostClient(cfg)
	postHandler := handler.NewPostHandler(postClient)

	chatClient := client.NewChatClient(cfg)
	chatHandler := handler.NewChatHandler(chatClient)

	serverHTTP := server.NewServerHTTP(authHandler, postHandler, chatHandler)
	return serverHTTP, nil
}
