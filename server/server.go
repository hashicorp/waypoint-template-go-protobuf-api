// Copyright (c) HashiCorp, Inc.

package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/%%gh_org%%/%%wp_project%%/config"
	%%wp_project%%v1 "github.com/%%gh_org%%/%%wp_project%%/gen/proto/go/%%wp_project%%/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
    _ "github.com/lib/pq"
)

type %%Wp_project%%Server struct {
	%%wp_project%%v1.Unimplemented%%Wp_project%%ServiceServer

	config config.%%Wp_project%%
}

// New%%Wp_project%%Server initializes a new server from config
func New%%Wp_project%%Server(config config.%%Wp_project%%) (*%%Wp_project%%Server, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := %%Wp_project%%Server{
		config: config,
	}

	return &server, nil
}

func (s * %%Wp_project%%Server) HelloWorld(
	ctx context.Context,
	req *%%wp_project%%v1.HelloWorldRequest,
) (*%%wp_project%%v1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &%%wp_project%%v1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}

func (s *%%Wp_project%%Server) ConnDB(
	ctx context.Context,
	req *%%wp_project%%v1.ConnDBRequest,
) (*%%wp_project%%v1.ConnDBResponse, error) {
	log.Printf("Connecting to Postgres")
	type conn struct {
		user     string
		password string
		port     int
		host     string
		dbname   string
	}

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Printf("Error converting DATABASE_PORT environment variable from string to int: %s", err.Error())
	}
	connection := conn{
		user:     os.Getenv("DATABASE_USERNAME"),
		password: os.Getenv("DATABASE_PASSWORD"),
		port:     port,
		host:     os.Getenv("DATABASE_HOSTNAME"),
		dbname:   os.Getenv("DATABASE_NAME"),
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connection.host, connection.port, connection.user, connection.password, connection.dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error connecting to Postgres: %s", err.Error())
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging Postgres DB: %s", err.Error())
		return nil, err
	}
	return &%%wp_project%%v1.ConnDBResponse{}, nil
}
