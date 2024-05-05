package web

import (
	"encoding/json"
	"net/http"

	"github.com/regismartiny/go-expert-desafio-cleanarchitecture/internal/entity"
	"github.com/regismartiny/go-expert-desafio-cleanarchitecture/internal/usecase"
	"github.com/regismartiny/go-expert-desafio-cleanarchitecture/pkg/events"
)

type WebOrderListHandler struct {
	EventDispatcher  events.EventDispatcherInterface
	OrderRepository  entity.OrderRepositoryInterface
	OrderListedEvent events.EventInterface
}

func NewWebOrderListHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderListedEvent events.EventInterface,
) *WebOrderListHandler {
	return &WebOrderListHandler{
		EventDispatcher:  EventDispatcher,
		OrderRepository:  OrderRepository,
		OrderListedEvent: OrderListedEvent,
	}
}

func (h *WebOrderListHandler) List(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository, h.OrderListedEvent, h.EventDispatcher)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
