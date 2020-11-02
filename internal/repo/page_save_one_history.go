package repo

import (
	"context"
)

func (repo *Repo) SaveOneHistory(paper *Page) (err error) {
	var (
		ctx context.Context
	)

	ctx = context.TODO()

	_, _ = repo.history.InsertOne(ctx, paper)

	return
}
