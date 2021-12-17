package interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		hasArea  float64
		name string
	}{
		{name: "rectangle", shape: Rectangle{width: 12, height: 6}, hasArea: 72},
		{name: "circle", shape:Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{base: 10, height: 5}, hasArea: 25},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T){
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf(" %#v got %g want %g", tt.shape,  got, tt.hasArea)
		}
	})
}
}
