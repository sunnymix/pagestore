package repo

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func (repo *Repo) FindOne(pid string) (page *Page, err error) {
	var (
		ctx    context.Context
		filter bson.M
	)

	ctx = context.TODO()
	filter = bson.M{"pid": pid}

	if err = repo.page.FindOne(ctx, filter).Decode(&page); err != nil {
		fmt.Printf("findone error: %s\n", err)
		return
	}
	return
}
