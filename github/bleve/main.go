package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

var index bleve.Index

func AddMsg() {
	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "23",
		From: "marty.schoch@gmail.com",
		Body: `协程(coroutine)顾名思义就是“协作的例程”（co-operative routines）。跟具有操作系统概念的线程不一样，协程是在用户空间利用程序语言的语法语义就能实现逻辑上类似多任务的编程技巧。实际上协程的概念比线程还要早，按照 Knuth 的说法“子例程是协程的特例”，一个子例程就是一次子函数调用，那么实际上协程就是类函数一样的程序组件，你可以在一个线程里面轻松创建数十万个协程，就像数十万次函数调用一样。只不过子例程只有一个调用入口起始点，返回之后就结束了，而协程入口既可以是起始点，又可以从上一个返回点继续执行，也就是说协程之间可以通过 yield 方式转移执行权，对称（symmetric）、平级地调用对方，而不是像例程那样上下级调用关系。当然 Knuth 的“特例”指的是协程也可以模拟例程那样实现上下级调用关系golang，这就叫非对称协程（asymmetric coroutines）。`,
	}
	index.Index(message.Id, message)
}

func Once() {
	var err error
	imp := bleve.NewIndexMapping()
	blogMapping := bleve.NewDocumentMapping()
	imp.AddDocumentMapping("blog", blogMapping)
	index, err = bleve.New("blogsdata", imp)
	if err != nil {
		panic(err)
	}
}

func Pre() {
	var err error
	index, err = bleve.Open("blogsdata")
	if err != nil {
		return
	}
}

func main() {
	Pre()
	AddMsg()
	query := bleve.NewMatchQuery("golang")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
