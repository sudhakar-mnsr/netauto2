package main

import (
   "fmt"
   "math"
)

type shape interface {
   area() float64
}

type polygon interface {
   shape
   perim()
}

type curved interface {
   shape
   circonf()
}

type rect struct {
   name string
   length,height float64
}
func (r *rect) area() float64 {
   return r.length * r.height
}
func (r *rect) perim() float64 { 
   return 2*r.length + 2*r.height 
} 
 
type triangle struct {
   name string
   a,b float64
} 
func (t *triangle) area() float64 { 
   return 0.5*(t.a * t.b) 
} 
func (t *triangle) perim() float64 { 
   return t.a + t.b + math.Sqrt((t.a*t.a) + (t.b*t.b)) 
}

type circle struct {
   name string
   rad float64
}
func (c *circle) area() float64 {
   return math.Pi * (c.rad*c.rad)
}
func (c *circle) circonf() float64 { 
   return 2 * math.Pi * c.rad 
}

func shapeInfo(s shape) string {
   return fmt.Sprintf("Area is %.2f", s.area())
}


func main() {
   r := &rect{"Rectangle", 4, 4}
   fmt.Println(r)
   fmt.Println(shapeInfo(r))
   c := &circle{"circle", 4}
   fmt.Println(c)
   fmt.Println(shapeInfo(c))
   fmt.Println(c.circonf())
}
