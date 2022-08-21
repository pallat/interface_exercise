package interfaceexercise

import "testing"

func TestCubeVolume(t *testing.T) {
	c := cube{edge: 2}
	if v := VolumeOf(c); v != 8 {
		t.Errorf("VolumeOf(cube) = %f, want 8", v)
	}
}

func TestTriangularPrismVolume(t *testing.T) {
	tP := triangularPrism{base: 2, attitude: 3, height: 4}
	if v := VolumeOf(tP); v != 12 {
		t.Errorf("VolumeOf(triangularPrism) = %f, want 12", v)
	}
}
