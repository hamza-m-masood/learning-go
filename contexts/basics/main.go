package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("doSomething: myKey's value is: ", ctx.Value("myKey"))
	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)
	fmt.Println("doSomething: myKey's value is: ", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))

}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myval")
	doSomething(ctx)
}
