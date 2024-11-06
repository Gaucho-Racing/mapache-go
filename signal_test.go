package mapache

import (
	"testing"
)

func scalingFunc(v float64) float64 {
	return v * 2
}

func TestSignal(t *testing.T) {
	apps := Signal{
		VehicleID:   "gr24",
		Name:        "pedal_pos1",
		RawValue:    10,
		ScalingFunc: scalingFunc,
	}
	apps.Scale()

	println(apps.Value) // 20
}
