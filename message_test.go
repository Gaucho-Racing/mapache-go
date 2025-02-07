package mapache

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestTimestamp(t *testing.T) {
	now := time.Now().UnixMicro()
	fmt.Println(now)
	fmt.Printf("Size of timestamp: %d bytes\n", unsafe.Sizeof(now))
}

func TestMessage(t *testing.T) {
	ecuStatusMessage := Message{
		NewField("ecu_state", 1, Unsigned, BigEndian, nil),
		NewField("ecu_status_flags", 3, Unsigned, BigEndian, func(f Field) []Signal {
			signals := []Signal{}
			bitMap := []string{
				"ecu_status_acu",
				"ecu_status_inv_one",
				"ecu_status_inv_two",
				"ecu_status_inv_three",
				"ecu_status_inv_four",
				"ecu_status_fan_one",
				"ecu_status_fan_two",
				"ecu_status_fan_three",
				"ecu_status_fan_four",
				"ecu_status_fan_five",
				"ecu_status_fan_six",
				"ecu_status_fan_seven",
				"ecu_status_fan_eight",
				"ecu_status_dash",
				"ecu_status_steering",
			}
			for i := 0; i < len(bitMap); i++ {
				signals = append(signals, Signal{
					Name:     bitMap[i],
					Value:    float64(f.CheckBit(i)),
					RawValue: f.CheckBit(i),
				})
			}
			return signals
		}),
		NewField("ecu_maps", 1, Unsigned, BigEndian, func(f Field) []Signal {
			signals := []Signal{}
			signals = append(signals, Signal{
				Name:     "ecu_power_level",
				Value:    float64((f.Value >> 4) & 0x0F),
				RawValue: (f.Value >> 4) & 0x0F,
			})
			signals = append(signals, Signal{
				Name:     "ecu_torque_map",
				Value:    float64(f.Value & 0x0F),
				RawValue: f.Value & 0x0F,
			})
			return signals
		}),
		NewField("ecu_max_cell_temp", 1, Unsigned, BigEndian, func(f Field) []Signal {
			signals := []Signal{}
			signals = append(signals, Signal{
				Name:     "ecu_max_cell_temp",
				Value:    float64(f.Value) * 0.25,
				RawValue: f.Value,
			})
			return signals
		}),
		NewField("ecu_acu_state_of_charge", 1, Unsigned, BigEndian, func(f Field) []Signal {
			signals := []Signal{}
			signals = append(signals, Signal{
				Name:     "ecu_acu_state_of_charge",
				Value:    float64(f.Value) * 20 / 51,
				RawValue: f.Value,
			})
			return signals
		}),
		NewField("ecu_glv_state_of_charge", 1, Unsigned, BigEndian, func(f Field) []Signal {
			signals := []Signal{}
			signals = append(signals, Signal{
				Name:     "ecu_glv_state_of_charge",
				Value:    float64(f.Value) * 20 / 51,
				RawValue: f.Value,
			})
			return signals
		}),
	}
	t.Run("Invalid byte length", func(t *testing.T) {
		err := ecuStatusMessage.FillFromBytes([]byte{0, 0})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
	t.Run("Valid byte length", func(t *testing.T) {
		err := ecuStatusMessage.FillFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		signals := ecuStatusMessage.ExportSignals()
		for _, signal := range signals {
			fmt.Printf("%s: %f\n", signal.Name, signal.Value)
		}
	})
}
