package main
import(
    "fmt"
    "io"
    "os"
    "io/ioutil"
    "bufio"
)

func main(){
    //readfile()    
    readfile2()
    readfile3()
}

func readfile(){
    path:="./xxx"
    fd,_:=os.Open(path)
    defer fd.Close()
    buf:=make([]byte,10)
    io.ReadFull(fd,buf)
    fmt.Printf("%v\n",buf)  //it only read 10 bytes
}
func readfile2(){
    path:="./xxx"
    fd,_:=os.Open(path)
    defer fd.Close()
    buf,_:=ioutil.ReadAll(fd)
    bufex,_:=ioutil.ReadFile(path)
    fmt.Printf("len of buf:%d,%d\n",len(buf),len(bufex))
}
func readfile3(){
    path:="./xxx"
    fd,_:=os.Open(path)
    defer fd.Close()
    r:=bufio.NewReader(fd)
    buf:=make([]byte,11)
    r.Read(buf)
    fmt.Printf("%v\n",buf)
}
