// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/event"
	"github.com/devfullcycle/20-CleanArch/internal/infra/database"
	"github.com/devfullcycle/20-CleanArch/internal/infra/web"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/google/wire"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	return createOrderUseCase
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	orderCreated := event.NewOrderCreated()
	webOrderHandler := web.NewWebOrderHandler(eventDispatcher, orderRepository, orderCreated)
	return webOrderHandler
}

func NewListOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	orderRepository := database.NewOrderRepository(db)
	orderListed := event.NewOrderListed()
	listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository, orderListed, eventDispatcher)
	return listOrdersUseCase
}

func NewWebOrderListHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderListHandler {
	orderRepository := database.NewOrderRepository(db)
	orderListed := event.NewOrderListed()
	webOrderListHandler := web.NewWebOrderListHandler(eventDispatcher, orderRepository, orderListed)
	return webOrderListHandler
}

// wire.go:

var setOrderRepositoryDependency = wire.NewSet(database.NewOrderRepository, wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setEventDispatcherDependency2 = wire.NewSet(events.NewEventDispatcher, event.NewOrderListed, wire.Bind(new(events.EventInterface), new(*event.OrderListed)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setOrderCreatedEvent = wire.NewSet(event.NewOrderCreated, wire.Bind(new(events.EventInterface), new(*event.OrderCreated)))

var setOrderListedEvent = wire.NewSet(event.NewOrderListed, wire.Bind(new(events.EventInterface), new(*event.OrderListed)))
