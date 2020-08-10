package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"paperstore/cfg"
	"time"
)

type Repo struct {
	client  *mongo.Client
	paper   *mongo.Collection
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
		paperCollection   *mongo.Collection
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

	paperCollection = client.Database("paper").Collection("paper")

	historyCollection = client.Database("paper").Collection("history")

	GlobalRepo = &Repo{
		client:  client,
		paper:   paperCollection,
		history: historyCollection,
	}
}

func init() {
	initRepo()
}
