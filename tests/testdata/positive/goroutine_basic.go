package main

func say(msg string) {
  fmt.Println(msg);
}

func main() {
  go say("from goroutine");
}
