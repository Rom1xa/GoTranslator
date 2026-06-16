package main

type Any interface {
}

func main() {
  var x Any = 10;
  fmt.Println(x);
  x = "text";
  fmt.Println(x);
}
