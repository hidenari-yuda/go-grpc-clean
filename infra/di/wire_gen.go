// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hidenari-yuda/go-grpc-clean/handler"
	"github.com/hidenari-yuda/go-grpc-clean/repository"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/hidenari-yuda/go-grpc-clean/usecase/interactor"
)

// Injectors from wire.go:

// User
func InitializeUserHandler(db repository.SQLExecuter, fb usecase.Firebase) handler.UserHandler {
	userRepository := repository.NewUserRepositoryImpl(db)
	userInteractor := interactor.NewUserInteractorImpl(fb, userRepository)
	userHandler := handler.NewUserHandlerImpl(userInteractor)
	return userHandler
}

// Chat
func InitializeChatHandler(db repository.SQLExecuter, fb usecase.Firebase) handler.ChatHandler {
	chatRepository := repository.NewChatRepositoryImpl(db)
	chatGroupRepository := repository.NewChatGroupRepositoryImpl(db)
	chatUserRepository := repository.NewChatUserRepositoryImpl(db)
	chatInteractor := interactor.NewChatInteractorImpl(fb, chatRepository, chatGroupRepository, chatUserRepository)
	chatHandler := handler.NewChatHandlerImpl(chatInteractor)
	return chatHandler
}

// User
func InitializeUserInteractor(db repository.SQLExecuter, fb usecase.Firebase) interactor.UserInteractor {
	userRepository := repository.NewUserRepositoryImpl(db)
	userInteractor := interactor.NewUserInteractorImpl(fb, userRepository)
	return userInteractor
}

// Chat
func InitializeChatInteractor(db repository.SQLExecuter, fb usecase.Firebase) interactor.ChatInteractor {
	chatRepository := repository.NewChatRepositoryImpl(db)
	chatGroupRepository := repository.NewChatGroupRepositoryImpl(db)
	chatUserRepository := repository.NewChatUserRepositoryImpl(db)
	chatInteractor := interactor.NewChatInteractorImpl(fb, chatRepository, chatGroupRepository, chatUserRepository)
	return chatInteractor
}

// wire.go:

var wireSet = wire.NewSet(handler.WireSet, interactor.WireSet, repository.WireSet)
