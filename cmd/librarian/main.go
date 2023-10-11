package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	librarianapi "github.com/dgdraganov/super-librarian"
	librarian "github.com/dgdraganov/super-librarian/gen/librarian"
	"github.com/dgdraganov/super-librarian/internal/model"
	"github.com/dgdraganov/super-librarian/internal/storage/repo"
)

func main() {

	// appConfig, err := config.NewServiceConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("new service config: %w", err))
	// }

	// dbConfig, err := config.NewDatabaseConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("new db config: %w", err))
	// }

	appConfig := model.ServiceConfig{
		AppEnv:   "DEBUG",
		HttpPort: "9205",
		Host:     "localhost",
	}
	dbConfig := model.MySqlConfig{
		Username: "docker",
		Password: "mysql_docker_password",
		Host:     "localhost",
		Port:     "3306",
		DbName:   "library",
	}

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[librarianapi] ", log.Ltime)
	}

	// Setup repository.
	var (
		repository librarianapi.BooksRepo
	)
	{
		repository = repo.NewLibraryRepository(dbConfig)
	}

	// Initialize the services.
	var (
		librarianSvc librarian.Service
	)
	{
		librarianSvc = librarianapi.NewLibrarian(repository, logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		librarianEndpoints *librarian.Endpoints
	)
	{
		librarianEndpoints = librarian.NewEndpoints(librarianSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch appConfig.Host {
	case "localhost":
		{
			addr := "http://0.0.0.0:8000"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s\n", addr, err)
			}
			// if *secureF {
			// 	u.Scheme = "https"
			// }
			// if *domainF != "" {
			// 	u.Host = *domainF
			// }
			if appConfig.HttpPort != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatalf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, appConfig.HttpPort)
				fmt.Println(u.Host)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			isProd := strings.ToUpper(appConfig.AppEnv) == "PROD"
			handleHTTPServer(ctx, u, librarianEndpoints, &wg, errc, logger, !isProd)
		}

	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: localhost)\n", appConfig.Host)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
