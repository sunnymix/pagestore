package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pagestore/cfg"
	"time"
)

type Repo struct {
	client  *mongo.Client
	page    *mongo.Collection
	history *mongo.Collection
}

var (
	GlobalRepo *Repo
)

func initRepo() {
	var (
		conn              *cfg.Conn
		ctx               context.Context
		opts              *options.ClientOptions
		client            *mongo.Client
		err               error
		pageCollection    *mongo.Collection
		historyCollection *mongo.Collection
	)

	conn, err = cfg.NewConn()
	if err != nil {
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	connStr := conn.ToString("mongodb")

	fmt.Printf("connStr: %s\n", connStr)

	opts = options.Client().ApplyURI(connStr)

	if client, err = mongo.Connect(ctx, opts); err != nil {
		fmt.Printf("connect error: %s\n", err)
		return
	}

	pageCollection = client.Database("page").Collection("page")

	historyCollection = client.Database("page").Collection("history")

	GlobalRepo = &Repo{
		client:  client,
		page:    pageCollection,
		history: historyCollection,
	}
}

func init() {
	initRepo()
}
