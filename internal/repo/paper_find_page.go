package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func (repo *Repo) FindPage(query string) (papers []*Paper, err error) {
	var (
		ctx        context.Context
		filter     bson.M
		projection bson.M
		opts       *options.FindOptions
	)

	ctx = context.TODO()

	filter = bson.M{}
	if len(query) > 0 {
		filter = bson.M{"title": bson.M{"$regex": query, "$options": "$i"}}
	}

	projection = bson.M{"pid": 1, "title": 1, "updated": 1, "_id": 0}

	size := int64(50)

	sort := bson.M{"title": 1}

	opts = &options.FindOptions{
		Projection: projection,
		Limit:      &size,
		Sort:       sort,
	}

	cursor, err := repo.paper.Find(ctx, filter, opts)

	if err != nil {
		return
	}

	_ = cursor.All(ctx, &papers)

	return
}
