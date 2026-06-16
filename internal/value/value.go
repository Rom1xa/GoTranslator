package value

import (
	"strconv"
	"strings"
)

// Value — значение времени выполнения
type Value interface {
	String() string
	Native() any
}

type Int int64

func (i Int) String() string { return strconv.FormatInt(int64(i), 10) }
func (i Int) Native() any    { return int64(i) }

type Float float64

func (f Float) String() string { return strconv.FormatFloat(float64(f), 'g', -1, 64) }
func (f Float) Native() any    { return float64(f) }

type String string

func (s String) String() string { return string(s) }
func (s String) Native() any    { return string(s) }

type Bool bool

func (b Bool) String() string { return strconv.FormatBool(bool(b)) }
func (b Bool) Native() any    { return bool(b) }

type Interface struct {
	TypeName string
	Value    Value
}

func NewInterface(typeName string, v Value) *Interface {
	return &Interface{TypeName: typeName, Value: v}
}

func (i *Interface) String() string {
	if i.Value == nil {
		return "<nil>"
	}
	return i.Value.String()
}

func (i *Interface) Native() any {
	if i.Value == nil {
		return nil
	}
	return i.Value.Native()
}

// Struct — значение типа struct.
type Struct struct {
	TypeName string
	fields   map[string]Value
}

func NewStruct(typeName string) *Struct {
	return &Struct{TypeName: typeName, fields: map[string]Value{}}
}

func (s *Struct) Get(name string) (Value, bool) {
	v, ok := s.fields[name]
	return v, ok
}

func (s *Struct) Set(name string, v Value) {
	s.fields[name] = v
}

func (s *Struct) String() string {
	result := "{"
	first := true
	for k, v := range s.fields {
		if !first {
			result += " "
		}
		result += k + ":" + v.String()
		first = false
	}
	return result + "}"
}

func (s *Struct) Native() any { return s }

// Slice — значение типа []ElemType.
type Slice struct {
	ElemType string
	elems    []Value
}

func NewSlice(elemType string) *Slice {
	return &Slice{ElemType: elemType, elems: []Value{}}
}

func (s *Slice) Len() int           { return len(s.elems) }
func (s *Slice) Get(i int) Value    { return s.elems[i] }
func (s *Slice) Set(i int, v Value) { s.elems[i] = v }

func (s *Slice) AppendInPlace(v Value) { s.elems = append(s.elems, v) }

func (s *Slice) Append(v Value) *Slice {
	n := make([]Value, len(s.elems)+1)
	copy(n, s.elems)
	n[len(s.elems)] = v
	return &Slice{ElemType: s.ElemType, elems: n}
}

func (s *Slice) String() string {
	parts := make([]string, len(s.elems))
	for i, e := range s.elems {
		parts[i] = e.String()
	}
	return "[" + strings.Join(parts, " ") + "]"
}

func (s *Slice) Native() any { return s }
