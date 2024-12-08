package stack

import "testing"

func TestTuringTape(t *testing.T) {
	tape := NewTuringTape(10)

	for i := 0; i < 10; i++ {
		tape.Write(i)
		_ = tape.MoveRight()
	}

	for i := 0; i < 10; i++ {
		value := tape.Look()
		if value != 9-i {
			t.Errorf("value is %d, but expected %d", value, 9-i)
		}

		_ = tape.MoveLeft()
	}
}
