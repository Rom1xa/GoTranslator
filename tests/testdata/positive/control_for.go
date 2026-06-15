package main
func main() {
  fact := 1;
  for i := 1; i <= 5; i = i + 1 {
    fact = fact * i;
  }
  fmt.Println(fact);
  k := 0;
  for k < 3 {
    fmt.Println(k);
    k = k + 1;
  }
}
