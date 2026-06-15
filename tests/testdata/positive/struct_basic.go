package main
type Point struct {
  x int;
  y int;
}
func main() {
  p := Point{x: 3, y: 4};
  fmt.Println(p.x);
  p.y = 10;
  fmt.Println(p.y);
  q := Point{x: 1};
  fmt.Println(q.x);
  fmt.Println(q.y);
}
