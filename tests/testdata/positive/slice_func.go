package main
func sum(nums []int) int {
  total := 0;
  for i := range nums {
    total = total + nums[i];
  }
  return total;
}
func push(nums []int, val int) []int {
  return append(nums, val);
}
func main() {
  s := []int{1, 2, 3, 4};
  fmt.Println(sum(s));
  s = push(s, 5);
  fmt.Println(len(s));
  fmt.Println(s[4]);
}
