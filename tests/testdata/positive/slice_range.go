package main
func main() {
  nums := []int{10, 20, 30};
  total := 0;
  for i, v := range nums {
    total = total + v;
    fmt.Printf("nums[%d] = %d\n", i, v);
  }
  fmt.Println(total);
}
