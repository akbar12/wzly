package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"pos-api/src/controller"
	"pos-api/src/repository"
	"pos-api/src/util"

	"github.com/gorilla/mux"
)

// import "pos_api/src/repository"

func main() {
	util.InitConfig()
	db := util.InitDB()
	invoiceRepo := repository.InitInvoiceRepo(db)
	invoiceItemRepo := repository.InitInvoiceItemRepo(db)
	itemRepo := repository.InitItemRepo(db)
	customerRepo := repository.InitCustomerRepo(db)
	invoiceController := controller.InitInvoiceCtrl(invoiceRepo, invoiceItemRepo, itemRepo, customerRepo)
	r := mux.NewRouter().StrictSlash(true)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(util.AuthMiddleware)

	r.HandleFunc("/v1/invoices/{id}", invoiceController.DetailInvoice).Methods("GET")
	r.HandleFunc("/v1/invoices/{id}", invoiceController.InvoiceUpdate).Methods("PUT")
	r.HandleFunc("/v1/invoices", invoiceController.InvoiceCreate).Methods("POST")
	r.HandleFunc("/v1/invoices", invoiceController.InvoiceList).Methods("GET")

	httpServer(r)
}

func httpServer(r *mux.Router) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	port := util.GetConfigString("http_server.port")
	srv := &http.Server{
		Addr: port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		fmt.Println("listening on port", port)
		if err := srv.ListenAndServe(); err != nil {
			util.Info(err, nil)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	util.Info("shutting down server", nil)
	os.Exit(0)
}
