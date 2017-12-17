package main
import(
	"fmt"
	"package/tools"
	"time"
	//bl "github.com/blake2b-master"
	//"crypto"
	"package/sha3"
	//"adler32"
)

const max=100*10000
var data = []byte{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
func testmd5(){
	ts :=time.Now().UnixNano()
	for i:=0;i<max;i++{
		tools.Md5Cal(data)
	}
	te := time.Now().UnixNano()
	tp :=(te-ts)/1000000
	fmt.Printf("testing md5,%d times cost:%d ms\n",max,tp)
}
func testsha1(){
	var val string
        ts :=time.Now().UnixNano()
        for i:=0;i<max;i++{
        	val=tools.Sha1Cal(data)
        }
        te := time.Now().UnixNano()
        tp :=(te-ts)/1000000
        fmt.Printf("testing sha1,%d times cost:%d ms,leng:%d\n",max,tp,len(val))
}
/*
func testBlake2(){
	ts :=time.Now().UnixNano()
        for i:=0;i<max;i++{
               v:=adler32.Checksum(data)
		fmt.Printf("val:%d\n",v)
        }       
        te := time.Now().UnixNano()
        tp :=(te-ts)/1000000
        fmt.Printf("testing blake2,%d times cost:%d ms\n",max,tp)
}
*/
func testsha3(){
	var val [32]byte
        ts :=time.Now().UnixNano()
        for i:=0;i<max;i++{
                val=sha3.Sum256(data)
        }
        te := time.Now().UnixNano()
        tp :=(te-ts)/1000000
        fmt.Printf("testing sha3,%d times cost:%d ms,leng:%d\n",max,tp,len(val))
}
func main(){
	//testmd5()	
	testsha1()	
	testsha3()	
}

