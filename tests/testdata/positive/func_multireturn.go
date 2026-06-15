package main
func minmax(a int, b int) (int, int) {
  if a < b {
    return a, b;
  }
  return b, a;
}
func main() {
  lo, hi := minmax(5, 3);
  fmt.Println(lo, hi);
  lo2, hi2 := minmax(1, 9);
  fmt.Println(lo2, hi2);
}
