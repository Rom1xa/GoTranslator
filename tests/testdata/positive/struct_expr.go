package main
type Rect struct {
  w int;
  h int;
}
func area(r Rect) int {
  return r.w * r.h;
}
func main() {
  r := Rect{w: 5, h: 3};
  fmt.Println(area(r));
  fmt.Println(r.w + r.h);
}
