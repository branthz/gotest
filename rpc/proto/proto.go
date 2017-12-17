package proto

import(
	"errors"
)

type Args struct{
        A,B int
}

type Param struct{
        Yu,Chu int
}

type Client struct{}

func (c *Client)Multiply(args *Args,answer *int) error{
        *answer=args.A*args.B
        return nil
}

func (c *Client)Divide(args *Args,q *Param) error{
        if args.B == 0 {
                return errors.New("divide by zero")
        }
        q.Yu=args.A/args.B
        q.Chu=args.A%args.B
        return nil
}
