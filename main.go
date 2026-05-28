package main

import (
	"context"
	"fmt"
	"study/sql"
)

func main() {
	ctx := context.Background()
	conn, err := sql.CreateConnection(ctx)
	if err != nil {
		fmt.Println(err)
	}
	if err := sql.CreateTables(ctx, conn); err != nil {
		fmt.Println(err)
	}
	if err := sql.InsertRow(*conn, ctx); err != nil {
		fmt.Println(err)
	}
	fmt.Println("строка успешно добалвена")
}
