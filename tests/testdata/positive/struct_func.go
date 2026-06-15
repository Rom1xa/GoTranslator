package main
type Vec2 struct {
  x int;
  y int;
}
func add(a Vec2, b Vec2) Vec2 {
  return Vec2{x: a.x + b.x, y: a.y + b.y};
}
func main() {
  u := Vec2{x: 1, y: 2};
  v := Vec2{x: 3, y: 4};
  w := add(u, v);
  fmt.Println(w.x);
  fmt.Println(w.y);
}
