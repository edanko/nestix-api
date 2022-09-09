package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	adapters2 "github.com/edanko/nestix-api/internal/adapters"
	"github.com/edanko/nestix-api/internal/config"
	"github.com/edanko/nestix-api/internal/domain/path"
	"github.com/edanko/nestix-api/internal/domain/sheetpathdet"
	"github.com/edanko/nestix-api/internal/service"
	"github.com/edanko/nestix-api/pkg/logs"
	"github.com/edanko/nestix-api/pkg/tenant"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	cfg := config.GetConfig()

	logger := logs.NewZerologLogger(cfg.Logger.Level)
	if cfg.App.Environment == "development" {
		// log.Logger = log.Logger.Level(zerolog.DebugLevel)
	}

	connMap := make(map[string]*sqlx.DB, 10)

	for name, t := range cfg.Tenants {
		conn, err := sqlx.ConnectContext(
			ctx,
			"sqlserver",
			fmt.Sprintf(
				"sqlserver://%s:%s@%s?database=%s",
				t.DB.User, t.DB.Password, t.DB.Host, t.DB.Database,
			),
		)
		if err != nil {
			log.Fatal().Err(err).Msg("error connecting to database")
		}
		connMap[name] = conn
	}
	defer func(connMap map[string]*sqlx.DB) {
		for _, conn := range connMap {
			conn.Close()
		}
	}(connMap)

	pathRepo := adapters2.NewPathRepository(connMap)
	sheetPathRepo := adapters2.NewSheetPathRepository(connMap)
	sheetPathDetRepo := adapters2.NewSheetPathDetRepository(connMap)
	orderRepo := adapters2.NewOrderRepository(connMap)
	productRepo := adapters2.NewProductRepository(connMap)
	visualRepo := adapters2.NewVisualRepository(connMap)
	machineRepo := adapters2.NewMachineRepository(connMap)
	inventoryRepo := adapters2.NewInventoryRepository(connMap)

	svc := service.New(
		pathRepo,
		sheetPathRepo,
		sheetPathDetRepo,
		orderRepo,
		productRepo,
		visualRepo,
		machineRepo,
		inventoryRepo,
	)

	svc.Something(ctx)

	ctx = tenant.ContextWithTenantID(ctx, "IBSV")

	paths, err := pathRepo.SearchByName(ctx, "10-562003")
	if err != nil {
		log.Fatal().Err(err).Msg("search by name failed")
	}
	pathsIDs := lo.Map[*path.Path, int64](paths, func(x *path.Path, _ int) int64 {
		return x.ID()
	})

	sheetpaths, _ := sheetPathRepo.GetByPathIDs(ctx, pathsIDs)

	var details [][]*sheetpathdet.Part

	for _, sheetpath := range sheetpaths {
		sheetpathdet, err := sheetPathDetRepo.ListPartsBySheetPathID(ctx, sheetpath.ID())
		if err != nil {
			log.Fatal().Err(err).Msg("4get failed")
		}
		details = append(details, sheetpathdet)
	}

	var pathid int64 = 16901 // 20631

	path, err := pathRepo.GetByID(ctx, pathid)
	if err != nil {
		logger.Fatal("2get failed", err, nil)
	}

	fmt.Println("nxname", path.Name())

	sheetpath, err := sheetPathRepo.GetByPathID(ctx, pathid)
	if err != nil {
		logger.Fatal("3get failed", err, nil)
	}

	sp, err := sheetPathDetRepo.ListPartsBySheetPathID(ctx, sheetpath.ID())
	if err != nil {
		logger.Fatal("4get failed", err, nil)
	}

	orderIDs := lo.Map[*sheetpathdet.Part, int64](sp, func(x *sheetpathdet.Part, _ int) int64 {
		return x.OrderID()
	})

	// spew.Dump(orderIDs)

	orders, _ := orderRepo.GetByIDs(ctx, orderIDs)
	_ = orders

	products, _ := productRepo.GetByIDs(ctx, orderIDs)
	_ = products

	for i := range orderIDs {
		n := sp[i]

		orderline := orders[i]

		fmt.Println("detail code", n.DetailCode())
		fmt.Println("order", orderline.OrderNo())
		fmt.Println("section", orderline.Section())
		// product, err := productRepo.GetByID(ctx, orderline.PartID.Int64)
		// if err != nil {
		// 	log.Fatal().Err(err).Msg("6get failed")
		// }
		product := products[i]
		fmt.Println("pos", product.PartNo.String)
		q := n.DetailCount() * orderline.Count()
		fmt.Println("quantity", q)

		w := n.Area() * *orderline.Thick() * product.Density.Float64
		fmt.Println("weight", w)
		fmt.Println("total weight", float64(q)*w)
		fmt.Println("quality", product.Quality.String)
		fmt.Println("len", product.Length.Float64)
		fmt.Println("wid", product.Width.Float64)

		fmt.Println("thickness", product.Thick.Float64)
		fmt.Println()
	}

	// mr := adapters.NewMasterRepository(cfg.Tenants["MR"].Master, cfg.Tenants["MR"].Site)
	// n, err := mr.ReadNest("15800")
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("read nest failed")
	// }
	//
	// spew.Dump(n)
}
