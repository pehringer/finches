package fungen

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
