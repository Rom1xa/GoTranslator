package main
func main() {
  s := []int{1, 2, 3};
  fmt.Println(len(s));
  fmt.Println(s[0]);
  s[2] = 99;
  fmt.Println(s[2]);
  s = append(s, 4);
  fmt.Println(len(s));
  fmt.Println(s[3]);
}
