package main

import (
	"fmt"
	"math"
)

//类都是结构体，但是结构体不一定是类

//结构体
//type Point struct {
//	X,Y float64
//}
//// GO函数，类似java的静态方法
//func Distence(p,q Point) float64 {
//	return math.Hypot(q.X-p.X,q.Y-p.Y)
//}

//func main() {
//	p := Point{1, 2}
//	q := Point{4, 6}
//
//	fmt.Print(Distence(p, q))
//}

//类,推荐用类
type Point2 struct {
	X, Y float64
}

// GO方法，类似java的实例方法，方法里包含了接受者(Receiver)-----(p Point2)
func (p Point2) Distence2(q Point2) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point2{1, 2}
	q := Point2{4, 6}

	fmt.Print(p.Distence2(q))
}
