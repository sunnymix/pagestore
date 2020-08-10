package repo

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func (repo *Repo) FindOne(pid string) (paper *Paper, err error) {
	var (
		ctx    context.Context
		filter bson.M
	)

	ctx = context.TODO()
	filter = bson.M{"pid": pid}

	if err = repo.paper.FindOne(ctx, filter).Decode(&paper); err != nil {
		fmt.Printf("findone error: %s\n", err)
		return
	}
	return
}
