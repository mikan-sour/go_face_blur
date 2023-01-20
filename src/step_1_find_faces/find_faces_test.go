package findfaces

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)
	defer func() {
		fmt.Println("tests complete!")
	}()
	go func() {
		// setup()
		code := m.Run()
		done <- true
		os.Exit(code)
	}()

	select {
	case <-timeout:
		fmt.Println("Test didn't finish in time")
		os.Exit(1)
	case <-done:
	}
}

func TestRandomFloat(t *testing.T) {
	i := 0
	for i < 30 {
		f := randomFloat()
		if f < RADIUS_MIN || f > RADIUS_MAX {
			t.Fatalf("randomFloat is supposed to be between %f and %f but got %f", RADIUS_MIN, RADIUS_MAX, f)
		}
		i++
	}
}

func TestSubTile(t *testing.T) {
	cases := []struct {
		minx, miny, spacing int
		maxXExp, maxYExp    int
	}{
		{
			minx: 1, miny: 1, spacing: 3,
			maxXExp: 4, maxYExp: 4,
		},
		{
			minx: 4, miny: 8, spacing: 12,
			maxXExp: 16, maxYExp: 20,
		},
		{
			minx: 1000, miny: 1000, spacing: 50,
			maxXExp: 1050, maxYExp: 1050,
		},
	}

	for _, tc := range cases {
		rect := subTile(tc.minx, tc.miny, tc.spacing)

		if rect.Min.X != tc.minx {
			t.Fatalf("expected Min.X %d but got %d", tc.minx, rect.Min.X)
		}

		if rect.Min.Y != tc.miny {
			t.Fatalf("expected Min.Y %d but got %d", tc.miny, rect.Min.Y)
		}

		if rect.Max.X != tc.maxXExp {
			t.Fatalf("expected Max.X %d but got %d", tc.maxXExp, rect.Max.X)
		}

		if rect.Max.Y != tc.maxYExp {
			t.Fatalf("expected Max.Y %d but got %d", tc.maxYExp, rect.Max.Y)
		}
	}
}
