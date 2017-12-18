package main
import(
	"reflect"
	"fmt"
)
func main(){
	var a ="10" 
	var tp string
	v:=reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		tp="float"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tp="int"
	case reflect.String :
		tp ="string"
	default :
		tp="unknown"
	}
	fmt.Println(tp)
}
