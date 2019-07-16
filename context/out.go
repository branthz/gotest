package main
import(
"fmt"
	"context"
	"time"
)

type result struct {
	err  error
	msg  string
}

type student struct{
	score int
	passed bool
}

func NewStudent()*student{
	s:=new(student)
	s.score=100
	s.passed=true
	return s
}

func overBoard() bool{
	return true
}

func (s *student)checkLocus(ctx context.Context) {
	for {
		select{
		case <-ctx.Done():
			return 
		default :
			re:=overBoard()	
			if re==true{
				s.score=0
				s.passed=false
				return	
			}
		}
	}
}

func (s *student)parking(ctx context.Context) error {
	var re = make(chan *result)
	go s.driving(re)
	go s.checkLocus(ctx) 
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

func (s *student)driving(re chan *result) {
	//too slow
	//time.Sleep(2 * 1e9)

	r := new(result)
	//running 
	time.Sleep(1 * 1e9)
	r.msg = "ok"
	re <- r
	return
}
