package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (repo *Repo) SaveOne(page *Page) (err error) {
	var (
		ctx    context.Context
		filter bson.M
		val    bson.M
		upsert bool
		opts   *options.UpdateOptions
	)

	ctx = context.TODO()
	filter = bson.M{"pid": page.Pid}

	page.Updated = time.Now().UTC().UnixNano() / 1000000
	val = bson.M{"$set": page}

	upsert = true
	opts = &options.UpdateOptions{
		Upsert: &upsert,
	}

	if _, err = repo.page.UpdateOne(ctx, filter, val, opts); err != nil {
		fmt.Printf("insert page error: %s\n", err)
		return
	}

	_ = repo.SaveOneHistory(page)

	return
}
