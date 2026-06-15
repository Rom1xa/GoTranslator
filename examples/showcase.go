//go:build ignore

package main

type Point struct {
  x int;
  y int;
}

func makePoint(x int, y int) Point {
  return Point{x: x, y: y};
}

func distance(a Point, b Point) int {
  dx := a.x - b.x;
  dy := a.y - b.y;
  if dx < 0 {
    dx = -dx;
  }
  if dy < 0 {
    dy = -dy;
  }
  return dx + dy;
}

func fib(n int) int {
  if n <= 1 {
    return n;
  }
  return fib(n-1) + fib(n-2);
}

func sumSlice(nums []int) int {
  total := 0;
  for i, v := range nums {
    total = total + v;
    _ := i;
  }
  return total;
}

func main() {
  // 1. Типы и арифметика
  var a int = 10;
  b := 3;
  fmt.Printf("%d + %d = %d\n", a, b, a+b);
  fmt.Printf("%d / %d = %d (остаток %d)\n", a, b, a/b, a%b);

  // 2. Строки и bool
  greeting := "Hello";
  lang := "GoSubset";
  fmt.Println(greeting + ", " + lang + "!");
  fmt.Println(a > b && b > 0);

  // 3. Условие
  x := 42;
  if x > 100 {
    fmt.Println("большое");
  } else if x > 10 {
    fmt.Println("среднее");
  } else {
    fmt.Println("маленькое");
  }

  // 4. Цикл for + break/continue
  sum := 0;
  for i := 1; i <= 10; i = i + 1 {
    if i == 5 {
      continue;
    }
    sum = sum + i;
    if sum > 30 {
      break;
    }
  }
  fmt.Println("sum =", sum);

  // 5. Рекурсия
  for n := 0; n <= 7; n = n + 1 {
    fmt.Printf("fib(%d) = %d\n", n, fib(n));
  }

  // 6. Структуры
  p1 := makePoint(0, 0);
  p2 := makePoint(3, 4);
  fmt.Printf("distance = %d\n", distance(p1, p2));
  p1.x = 1;
  fmt.Printf("p1.x = %d\n", p1.x);

  // 7. Срезы
  nums := []int{};
  for i := 1; i <= 5; i = i + 1 {
    nums = append(nums, i*i);
  }
  fmt.Println("len =", len(nums));
  fmt.Printf("nums[4] = %d\n", nums[4]);
  nums[0] = 100;
  fmt.Println("sum of squares =", sumSlice(nums));

  // 8. Срез структур
  points := []Point{};
  points = append(points, makePoint(1, 2));
  points = append(points, makePoint(3, 4));
  for i, p := range points {
    fmt.Printf("points[%d] = (%d, %d)\n", i, p.x, p.y);
  }
}
