package container

import (
	"github.com/go-playground/validator"

	"anymind/pkg/infrastructure/gorm"
	"anymind/pkg/shared/config"
	Database "anymind/pkg/shared/database"
	historySvc "anymind/pkg/usecase/history"
	transactionSvc "anymind/pkg/usecase/transactions"
)

type Container struct {
	Config         *config.Config
	HistorySvc     historySvc.Service
	TransactionSvc transactionSvc.Service

	Validate *validator.Validate
}

func Setup() *Container {
	// ====== Construct Config
	cfg := config.NewConfig("./resources/config.json")

	// ====== Construct Database
	dbMaster, err := Database.New(cfg.Database.Master)
	if err != nil {
		panic(err)
	}
	dbSlave, _ := Database.New(cfg.Database.Slave)
	if dbSlave == nil {
		dbSlave = dbMaster
	}

	historyRepo := gorm.HistorySetup(dbMaster, dbSlave)
	accountRepo := gorm.AccountSetup(dbMaster, dbSlave)
	transactionRepo := gorm.TransactionsSetup(dbMaster, dbSlave)

	historySvc := historySvc.NewService(historyRepo)
	transactionSvc := transactionSvc.NewService(transactionRepo, accountRepo, historySvc)

	return &Container{
		Config:         cfg,
		HistorySvc:     historySvc,
		TransactionSvc: transactionSvc,
		Validate:       validator.New(),
	}
}
