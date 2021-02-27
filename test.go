package main

import (
	"fmt"
	"time"
)

//148 131 115

//func main(){
//	var a []uint8
//	var b string
//	a = []uint8{148,131,15}
//	c := ""
//	for i :=0; i<3; i++ {
//		n := a[i]
//		x := int64(n)
//		b = dectohex(x)
//		c += strings.Join([]string{b}, "")
//	}
//	fmt.Println(c)
//}
//
//
//
//func DecHex(dec []uint8) string {
//	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70} //声明并初始化一个key和value都是int64的map
//	s := ""
//	for i :=0; i<3; i++ {
//		n := dec[i]
//		x := int64(n)
//		fmt.Println(x)
//		for q := x; q > 0; q = q / 16 {
//			m := q % 16
//			if m > 9 && m < 16 {
//				m = hex[m]
//				s = fmt.Sprintf("%v%v", string(m), s)
//				continue
//			}
//			s = fmt.Sprintf("%v%v", m, s)
//		}
//	}
//	fmt.Println(s)
//	if len(s) == 1{
//		s = "0" + s
//	}
//	return s
//}
//
//func dectohex(n int64) string {
//	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
//	s := ""
//	for q := n; q > 0; q = q / 16 {
//		m := q % 16
//		if m > 9 && m < 16 {
//			m = hex[m]
//			fmt.Println(m)
//			s = fmt.Sprintf("%v%v", string(m), s)
//			continue
//		}
//		fmt.Println(m)
//		s = fmt.Sprintf("%v%v", m, s)
//	}
//	if len(s) == 1{
//		s = "0" + s
//	}
//	return s
//}
//var n = "123";
//
//func A(m *string){
//	fmt.Println(m)
//	n = "456"
//	fmt.Println(n)
//}
//
//func main()  {
//	A(&n)
//	fmt.Println(n)
//}
/////////////////////////////////////////////
//func init() {
//	// Log as JSON instead of the default ASCII formatter.
//	log.SetFormatter(&log.JSONFormatter{})
//
//	// Output to stdout instead of the default stderr
//	log.SetOutput(os.Stdout)
//
//	// Only log the warning severity or above.
//	log.SetLevel(log.InfoLevel)
//}
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	name := r.URL.Query().Get("name")
//	if name == "" {
//		name = "World"
//	}
//	log.Info("Received request for ", name)
//	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
//}
//
//func main() {
//	r := mux.NewRouter()
//	r.HandleFunc("/", handler)
//
//	srv := &http.Server{
//		Handler:      r,
//		Addr:         ":8080",
//	}
//	log.Info("Starting Server...")
//	if err := srv.ListenAndServe(); err != nil {
//		log.Fatal(err)
//	}
//}
//////////////////////////////////////////////////////

//将参数传入匿名函数 这样参数就可以压栈了
func goRun() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range values {
		go func(v int) {
			fmt.Println(v)
		}(v)
	}
}

//将参数赋值给临时变量 也可以将参数压栈
func goRun3() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range values {
		tmp := v
		go func() {
			fmt.Println(tmp)
		}()
	}
}

func main() {
	goRun3()
	time.Sleep(time.Second)
}
