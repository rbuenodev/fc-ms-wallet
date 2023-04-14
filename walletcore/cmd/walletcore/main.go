package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rbueno/fc-ms-wallet/internal/database"
	"github.com/rbueno/fc-ms-wallet/internal/event"
	createaccount "github.com/rbueno/fc-ms-wallet/internal/usecase/createAccount"
	createclient "github.com/rbueno/fc-ms-wallet/internal/usecase/createClient"
	createtransaction "github.com/rbueno/fc-ms-wallet/internal/usecase/createTransaction"
	"github.com/rbueno/fc-ms-wallet/pkg/events"
	"github.com/rbueno/fc-ms-wallet/web"
	"github.com/rbueno/fc-ms-wallet/web/webserver"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)
	transactionDB := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDB, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDB, accountDB, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3336")

	clientHandler := web.NewClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)
	webserver.Start()
}
