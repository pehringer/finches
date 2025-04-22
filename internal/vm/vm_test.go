package vm

import (
	"reflect"
	"testing"
)

func TestSetRegisters0(t *testing.T) {
	imm := []float32{}
	m := SetRegisters(imm)
	if m != nil {
		t.Errorf("expected: %v, actual: %v", nil, m)
	}
}

func TestSetRegisters1(t *testing.T) {
	imm := []float32{42}
	m := SetRegisters(imm)
	vm := &Machine{0, []float32{42}}
	if !reflect.DeepEqual(vm, m) {
		t.Errorf("expected: %v, actual: %v", vm, m)
	}
	reg := m.GetRegisters()
	if  !reflect.DeepEqual(imm, reg) {
		t.Errorf("expected: %v, actual: %v", imm, reg)
	}
}

func TestSetRegisters256(t *testing.T) {
	imm := make([]float32, 256)
	m := SetRegisters(imm)
	vm := &Machine{0, make([]float32, 256)}
	if !reflect.DeepEqual(vm, m) {
		t.Errorf("expected: %v, actual: %v", vm, m)
	}
	reg := m.GetRegisters()
	if !reflect.DeepEqual(imm, reg) {
		t.Errorf("expected: %v, actual: %v", imm, reg)
	}
}

func TestSetRegisters257(t *testing.T) {
	imm := make([]float32, 257)
	if m := SetRegisters(imm); m != nil {
		t.Errorf("expected: %v, actual: %v", nil, m)
	}
}

func TestOperationADD(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x00010203)
	m.Execute(0x00040506)
	m.Execute(0x00070809)
	reg := m.GetRegisters()
	exp := []float32{2, 8, 10, 0, 10}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestOperationSUB(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x02010203)
	m.Execute(0x02040506)
	m.Execute(0x02070809)
	reg := m.GetRegisters()
	exp := []float32{2, 8, 6, 0, -6}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestOperationMUL(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x04010203)
	m.Execute(0x04040506)
	m.Execute(0x04070809)
	reg := m.GetRegisters()
	exp := []float32{2, 0, 0, 0, 0}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestOperationDIV(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x06010203)
	m.Execute(0x06040506)
	m.Execute(0x06070809)
	reg := m.GetRegisters()
	exp := []float32{2, 8, 0, 0, 0.25}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestOperationMAX(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x08010203)
	m.Execute(0x08040506)
	m.Execute(0x08070809)
	reg := m.GetRegisters()
	exp := []float32{2, 8, 8, 0, 8}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestOperationMIN(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x0A010203)
	m.Execute(0x0A040506)
	m.Execute(0x0A070809)
	reg := m.GetRegisters()
	exp := []float32{2, 0, 0, 0, 0}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestOperationABD(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x0C010203)
	m.Execute(0x0C040506)
	m.Execute(0x0C070809)
	reg := m.GetRegisters()
	exp := []float32{2, 8, 6, 0, 6}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestOperationAVG(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x0E010203)
	m.Execute(0x0E040506)
	m.Execute(0x0E070809)
	reg := m.GetRegisters()
	exp := []float32{2, 4, 1.5, 0, 3}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestNoOperations(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x10000102)
	m.Execute(0x12010203)
	m.Execute(0x14020304)
	m.Execute(0x16030405)
	m.Execute(0x18040506)
	m.Execute(0x1A050607)
	m.Execute(0x1C060708)
	m.Execute(0x1E070809)
	m.Execute(0x20000102)
	m.Execute(0x22010203)
	m.Execute(0x24020304)
	m.Execute(0x26030405)
	m.Execute(0x28040506)
	m.Execute(0x2A050607)
	m.Execute(0x2C060708)
	m.Execute(0x2E070809)
	m.Execute(0x30000102)
	m.Execute(0x32010203)
	m.Execute(0x34020304)
	m.Execute(0x36030405)
	m.Execute(0x38040506)
	m.Execute(0x3A050607)
	m.Execute(0x3C060708)
	m.Execute(0x3E070809)
	reg := m.GetRegisters()
	exp := []float32{2, 4, 8, 0, 1}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestConditionLT(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x03000102)
	m.Execute(0x42010203)
	m.Execute(0x82020304)
	m.Execute(0xC2030405)
	m.Execute(0x42040506)
	m.Execute(0x82050607)
	m.Execute(0xC2060708)
	reg := m.GetRegisters()
	exp := []float32{-4, 8, 8, 0, -12}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestConditionGT(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x01000102)
	m.Execute(0x40010203)
	m.Execute(0x80020304)
	m.Execute(0xC0030405)
	m.Execute(0x40040506)
	m.Execute(0x80050607)
	m.Execute(0xC0060708)
	reg := m.GetRegisters()
	exp := []float32{5, 4, 1, 0, 1}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestConditionEQ(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x05010203)
	m.Execute(0x44020304)
	m.Execute(0x84030405)
	m.Execute(0xC4040506)
	m.Execute(0x44050607)
	m.Execute(0x84060708)
	m.Execute(0xC4070809)
	reg := m.GetRegisters()
	exp := []float32{2, 0, 0, 0, 0}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}

func TestSetFlags(t *testing.T) {
	imm := []float32{2, 4, 8, 0, 1}
	m := SetRegisters(imm)
	m.Execute(0x03000102)
	m.Execute(0x45010203)
        m.Execute(0xC1020304)
	m.Execute(0x86030405)
	reg := m.GetRegisters()
	exp := []float32{-4, 0, 1, -0.25, 1}
	if !reflect.DeepEqual(exp, reg) {
		t.Errorf("expected: %v, actual: %v", exp, reg)
	}
}
