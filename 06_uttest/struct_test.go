package main

import (
	"math"
	"testing"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

// 如下是直接使用interface实现，上层可以通过接口调用不同结构体的方法，却不用感知结构体之间的差别
// func TestArea(t *testing.T) {

// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %.2f, want %.2f", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		checkArea(t, circle, math.Pi*100.0)
// 	})

// }

// testForBranch

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

// func Area(r Rectangle) float64 {
// 	return r.width * r.height
// }

func main() {
	testing.Main(nil, nil, nil, nil)
	// 上述代码中的 testing.Main 函数将自动运行所有名为 Test* 的测试函数。
	// 不需要显式调用 testing.RunTests。
	//运行测试时，你可以使用 go test 命令，它会自动查找和运行测试文件。
}
