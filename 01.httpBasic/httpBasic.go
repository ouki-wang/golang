package main

import(
	"fmt"
	"reflect"
)
type HHH func(string, *string)

type HandlerFunc func(string, *string)

func (f HandlerFunc) ServeHTTP(w string, r *string) {
	f(w, r)
}

func x(name string, city *string){
	fmt.Println("name="+name);
	fmt.Println("city="+*city);
}
type Handler interface{
	ServeHTTP(w string, r *string)
}
func main(){
	var h Handler;
	type Host struct{
		Handler
	}
	y := HHH(x)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
	city := "beijing";
	h = HandlerFunc(x)
	h.ServeHTTP("wanghui",&city)
}
