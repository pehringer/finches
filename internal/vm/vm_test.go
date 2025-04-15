package vm

import (
	"errors"
	"reflect"
	"testing"
)

func TestInputOutputNone(t *testing.T) {
	in := []float32{}
	if out := SetInputs(in).GetOutputs(); !reflect.DeepEqual(in, out) {
		t.Errorf("incorrect outputs")
	}
}

func TestInputOutputSome(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	if out := SetInputs(in).GetOutputs(); !reflect.DeepEqual(in, out) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsADD(t *testing.T) {
	in := []float32{0}
	m := SetInputs(in)
	op := [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestADD(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x00, 0x40, 0x40, 0x40, 0x40}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{12}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsADDI(t *testing.T) {
	in := []float32{}
	m := SetInputs(in)
	op := [5]byte{0x02, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestADDI(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x02, 0x40, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x02, 0x40, 0x40, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x02, 0x40, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x02, 0x3F, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x02, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 3, 15}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsSUB(t *testing.T) {
	in := []float32{0}
	m := SetInputs(in)
	op := [5]byte{0x04, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestSUB(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x04, 0x40, 0x40, 0x40, 0x40}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{-4}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsSUBI(t *testing.T) {
	in := []float32{}
	m := SetInputs(in)
	op := [5]byte{0x06, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}


func TestSUBI(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x06, 0x40, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x06, 0x40, 0x40, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x06, 0x40, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x06, 0x3F, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x06, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 3, -5}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsMUL(t *testing.T) {
	in := []float32{0}
	m := SetInputs(in)
	op := [5]byte{0x08, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestMUL(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x08, 0x40, 0x40, 0x40, 0x40}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsMULI(t *testing.T) {
	in := []float32{}
	m := SetInputs(in)
	op := [5]byte{0x0A, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestMULI(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x0A, 0x40, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x0A, 0x40, 0x40, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x0A, 0x40, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x0A, 0x3F, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x0A, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 3, 0}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsDIV(t *testing.T) {
	in := []float32{0}
	m := SetInputs(in)
	op := [5]byte{0x0C, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
	in = []float32{1, 0}
	m = SetInputs(in)
	if err := m.Execute(op); !errors.Is(err, ErrorDivideByZero) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestDIV(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x0C, 0x40, 0x40, 0x40, 0x40}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsDIVI(t *testing.T) {
	in := []float32{}
	m := SetInputs(in)
	op := [5]byte{0x0E, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
	in = []float32{1}
	m = SetInputs(in)
	if err := m.Execute(op); !errors.Is(err, ErrorDivideByZero) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestDIVI(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x0E, 0x40, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x0E, 0x40, 0x40, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x0E, 0x40, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x0E, 0x3F, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x0E, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorDivideByZero) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 3, 0.20833333}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsMAX(t *testing.T) {
	in := []float32{0}
	m := SetInputs(in)
	op := [5]byte{0x10, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestMAX(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x10, 0x40, 0x40, 0x40, 0x40}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{5}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsMAXI(t *testing.T) {
	in := []float32{}
	m := SetInputs(in)
	op := [5]byte{0x12, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestMAXI(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x12, 0x40, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x12, 0x40, 0x40, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x12, 0x40, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x12, 0x3F, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x12, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 3, 5}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsMIN(t *testing.T) {
	in := []float32{0}
	m := SetInputs(in)
	op := [5]byte{0x14, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestMIN(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x14, 0x40, 0x40, 0x40, 0x40}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsMINI(t *testing.T) {
	in := []float32{}
	m := SetInputs(in)
	op := [5]byte{0x16, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestMINI(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x16, 0x40, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x16, 0x40, 0x40, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x16, 0x40, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x16, 0x3F, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x16, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 3, 0}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestPUSH(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x1A, 0x40, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x1A, 0x40, 0x40, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x1A, 0x40, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x1A, 0x3F, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x1A, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 3, 5, 4, 3, 2, 1, 0}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsSWP(t *testing.T) {
	in := []float32{0}
	m := SetInputs(in)
	op := [5]byte{0x1C, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorStackUnderflow) {
		t.Errorf("incorrect error")
	}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestSWP(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x1C, 0x40, 0x40, 0x40, 0x40}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 5, 3}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestErrorsPICK(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x1E, 0xBF, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorIndexOutOfBounds) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x1E, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorIndexOutOfBounds) {
		t.Errorf("incorrect error")
	}

	op = [5]byte{0x1E, 0x40, 0xE0, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, ErrorIndexOutOfBounds) {
		t.Errorf("incorrect error")
	}

	if out := m.GetOutputs(); !reflect.DeepEqual(out, in) {
		t.Errorf("incorrect outputs")
	}
}

func TestPICK(t *testing.T) {
	in := []float32{0, 1, 1, 2, 3, 5}
	m := SetInputs(in)
	op := [5]byte{0x1E, 0x00, 0x00, 0x00, 0x04}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x1E, 0x00, 0x00, 0x00, 0x03}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x1E, 0x00, 0x00, 0x00, 0x02}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 1, 1, 2, 3, 5, 1, 3, 1}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestFlagsNegative(t *testing.T) {
	in := []float32{0, -42}
	m := SetInputs(in)
	op := [5]byte{0x01, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x3A, 0x3F, 0x80, 0x00, 0x00} //LT PUSH 1
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x5A, 0x40, 0x00, 0x00, 0x00} //LE PUSH 2
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x7A, 0x40, 0x40, 0x00, 0x00} //EQ PUSH 3
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x9A, 0x40, 0x80, 0x00, 0x00} //NE PUSH 4
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xBA, 0x40, 0xA0, 0x00, 0x00} //GE PUSH 5
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xDA, 0x40, 0xC0, 0x00, 0x00} //GT PUSH 6
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xFA, 0x40, 0xE0, 0x00, 0x00} //NV PUSH 7
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{-42, 1, 2, 4}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestFlagsZero(t *testing.T) {
	in := []float32{0, 0}
	m := SetInputs(in)
	op := [5]byte{0x01, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x3A, 0x3F, 0x80, 0x00, 0x00} //LT PUSH 1
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x5A, 0x40, 0x00, 0x00, 0x00} //LE PUSH 2
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x7A, 0x40, 0x40, 0x00, 0x00} //EQ PUSH 3
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x9A, 0x40, 0x80, 0x00, 0x00} //NE PUSH 4
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xBA, 0x40, 0xA0, 0x00, 0x00} //GE PUSH 5
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xDA, 0x40, 0xC0, 0x00, 0x00} //GT PUSH 6
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xFA, 0x40, 0xE0, 0x00, 0x00} //NV PUSH 7
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{0, 2, 3, 5}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}

func TestFlagsPositive(t *testing.T) {
	in := []float32{0, 42}
	m := SetInputs(in)
	op := [5]byte{0x01, 0x00, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x3A, 0x3F, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x5A, 0x40, 0x00, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x7A, 0x40, 0x40, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0x9A, 0x40, 0x80, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xBA, 0x40, 0xA0, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xDA, 0x40, 0xC0, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	op = [5]byte{0xFA, 0x40, 0xE0, 0x00, 0x00}
	if err := m.Execute(op); !errors.Is(err, nil) {
		t.Errorf("incorrect error")
	}
	exp := []float32{42, 4, 5, 6}
	if out := m.GetOutputs(); !reflect.DeepEqual(out, exp) {
		t.Errorf("incorrect outputs")
	}
}
