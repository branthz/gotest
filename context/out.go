package main
import(
"fmt"
	"context"
	"time"
)
type result struct{
	msg string 
	code int
}

type client struct{
	resp *result 
	err error
	ctx context.Context
}

func newClient() (*client,context.Cancel){
	c:=new(client)
	ctx,cancel:=context.WithCancel(context.Background())
	c.ctx=ctx
	return c,cancel
}
