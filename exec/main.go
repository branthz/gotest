package main
import(
	"os/exec"
	"fmt"
	"bytes"
)

func main(){
	cmd := exec.Command("/bin/cp","-r","/home/brant/temp/openab", "/home/brant/temp/debian")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
    		fmt.Println(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
