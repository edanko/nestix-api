package service

import (
	"context"

	adapters2 "github.com/edanko/nestix-api/internal/adapters"
)

type Service struct {
	pathRepository         *adapters2.PathRepository
	sheetPathRepository    *adapters2.SheetPathRepository
	sheetPathDetRepository *adapters2.SheetPathDetRepository
	orderRepository        *adapters2.OrderRepository
	productRepository      *adapters2.ProductRepository
	visualRepository       *adapters2.VisualRepository
	machineRepository      *adapters2.MachineRepository
	inventoryRepository    *adapters2.InventoryRepository
}

func New(
	pathRepo *adapters2.PathRepository,
	sheetPathRepo *adapters2.SheetPathRepository,
	sheetPathDetRepo *adapters2.SheetPathDetRepository,
	orderRepo *adapters2.OrderRepository,
	productRepo *adapters2.ProductRepository,
	visualRepo *adapters2.VisualRepository,
	machineRepo *adapters2.MachineRepository,
	inventoryRepo *adapters2.InventoryRepository,
) *Service {
	return &Service{
		pathRepository:         pathRepo,
		sheetPathRepository:    sheetPathRepo,
		sheetPathDetRepository: sheetPathDetRepo,
		orderRepository:        orderRepo,
		productRepository:      productRepo,
		visualRepository:       visualRepo,
		machineRepository:      machineRepo,
		inventoryRepository:    inventoryRepo,
	}
}

func (s Service) Something(ctx context.Context) {

}
