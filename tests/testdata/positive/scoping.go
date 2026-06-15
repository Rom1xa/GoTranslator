package main
func main() {
  x := 10;
  for i := 0; i < 2; i = i + 1 {
    x := i;
    fmt.Println(x);
  }
  fmt.Println(x);
}
