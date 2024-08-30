package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	http2 "hex_ddd_cqs_example/http"
	"hex_ddd_cqs_example/internal/config"
	"hex_ddd_cqs_example/internal/domain/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.PostgresUser(),
		conf.PostgresPassword(),
		conf.PostgresHost(),
		conf.PostgresPort(),
		conf.PostgresDB(),
	))
	if err != nil {
		log.Fatal(err)
		return
	}

	sess, closeFunc, err := NewPostgreSQLConnection(conn)()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() { _ = closeFunc }()

	ctx, cancel := context.WithCancel(context.Background())

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())

	userRepo := user.NewPostgresUserRepository(sess) //NEW Declaring user repository

	r.GET("/", http2.Homepage())

	// NEW, changing the implementation in http controllers to use
	// user repository instead of database session
	r.POST("/users", http2.CreteUser(userRepo))
	r.GET("/users/:id", http2.FindUser(userRepo))

	server := &http.Server{
		Addr:    conf.ApiHost(),
		Handler: r,
	}
	defer func() { _ = server.Shutdown(ctx) }()

	fmt.Println("system start")
	defer func() {
		fmt.Println("system stop")
	}()

	readyCh := make(chan struct{})
	go func() {
		fmt.Println("system ready to serve")
		<-readyCh
	}()

	/* signal handling */
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		fmt.Println("signal trapped")
		cancel()
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func NewPostgreSQLConnection(conn *sql.DB) OpenDbConnectionFunc {
	return func() (db.Session, CloseDBFunc, error) {
		sess, err := postgresql.New(conn)
		if err != nil {
			return nil, nil, err
		}
		return sess, func() error { return conn.Close() }, nil
	}
}

type OpenDbConnectionFunc func() (db.Session, CloseDBFunc, error)

type CloseDBFunc func() error
