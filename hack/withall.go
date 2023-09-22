package main

import (
	"context"
	"gdb-withall-bug/internal/dao"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"log"
)

func main() {
	ctx := context.Background()
	post, err := dao.Post.Get(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v is not equals to %v", post.CreatedAt.Unix(), post.User.CreatedAt.Unix())
}
