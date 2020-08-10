package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (repo *Repo) SaveOne(paper *Paper) (err error) {
	var (
		ctx    context.Context
		filter bson.M
		val    bson.M
		upsert bool
		opts   *options.UpdateOptions
	)

	ctx = context.TODO()
	filter = bson.M{"pid": paper.Pid}

	paper.Updated = time.Now().UTC().UnixNano() / 1000000
	val = bson.M{"$set": paper}

	upsert = true
	opts = &options.UpdateOptions{
		Upsert: &upsert,
	}

	if _, err = repo.paper.UpdateOne(ctx, filter, val, opts); err != nil {
		fmt.Printf("insert paper error: %s\n", err)
		return
	}

	_ = repo.SaveOneHistory(paper)

	return
}
