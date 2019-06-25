package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background(), time.Second)
	defer cancel()

	go zcall(new(resourceA), ctx)

	go zcall(new(resourceB), ctx)

	go zcall(new(resourceB), ctx)

}

func zcall(x call, ctx context.Context) error {
	var r = make(chan *result)
	x.Rpc(ctx, r)
	select {
	case <-ctx.Done():
		r.err = ctx.Err()
	case x := <-r:
		if x.err != nil {

		}
	}
}

type call interface {
	Rpc(ctx context.Context, r chan *result)
}

type resourceA struct{}
type resourceB struct{}
type resourceC struct{}

func (c *resourceA) Rpc(ctx context.Context, resp chan *result) {
	time.Sleep(1e9)
	r = new(result)
	r.msg = "ok"
	resp <- r
	return
}

func (c *resourceB) Rpc(ctx context.Context, resp chan *result) {
	//...
	time.Sleep(1e9)
	r = new(result)
	r.msg = "ok"
	resp <- r
	return
}

func (c *resourceC) Rpc(ctx context.Context, resp chan *result) {
	r = new(result)
	r.msg = "falied"
	r.err = fmt.Errorf("Dial to 6.6.6.6 failed")
	resp <- r
	return
}

type result struct {
	err  error
	msg  string
	data interface{}
}

func parking(ctx context.Context) error {
	var re = make(chan *result)
	go driving(re)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case x := <-re:
		if x.err != nil {
			return x.err
		}
		fmt.Printf("parkging test:%s\n", x.msg)
		return nil
	}
}

func driving(re chan *result) {
	//something hanppend
	time.Sleep(2 * 1e9)
	r := new(result)
	r.msg = "ok"
	re <- r
	return
}
