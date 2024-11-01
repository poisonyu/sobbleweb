package main

import (
	"bytes"
	"fmt"
	"go/scanner"
	"log"
	"reflect"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

var strangelove = Movie{
	Title:    "Dr. Strangelove",
	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	Year:     1964,
	Color:    false,
	Actor: map[string]string{
		"Dr. Strangelove":            "Peter Sellers",
		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		"Pres. Merkin Muffley":       "Peter Sellers",
		"Gen. Buck Turgidson":        "George C. Scott",
		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	},

	Oscars: []string{
		"Best Actor (Nomin.)",
		"Best Adapted Screenplay (Nomin.)",
		"Best Director (Nomin.)",
		"Best Picture (Nomin.)",
	},
}

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return "lex.scan."
}
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())
	case reflect.Complex64, reflect.Complex128: // 1+2i => #C(1.0 2.0)
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%.1f %.1f)", real(c), imag(c))
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			fmt.Fprint(buf, "nil")
		}
	// case reflect.Ptr:
	case reflect.Pointer:
		return encode(buf, v.Elem())
	case reflect.Interface: // ("[]int"(1 2 3))
		buf.WriteByte('(')
		fmt.Fprintf(buf, "%q", v.Elem().Type())
		if err := encode(buf, v.Elem()); err != nil {
			return err
		}
		buf.WriteByte(')')
	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte('\n')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	default: // float, chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// func encode(indent int, buf *bytes.Buffer, v reflect.Value) error {
// 	switch v.Kind() {
// 	case reflect.Invalid:
// 		buf.WriteString("nil")
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 		fmt.Fprintf(buf, "%d", v.Int())
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 		fmt.Fprintf(buf, "%d", v.Uint())
// 	case reflect.Float32, reflect.Float64:
// 		fmt.Fprintf(buf, "%f", v.Float())
// 	case reflect.Complex64, reflect.Complex128: // 1+2i => #C(1.0 2.0)
// 		c := v.Complex()
// 		fmt.Fprintf(buf, "#C(%.1f %.1f)", real(c), imag(c))
// 	case reflect.String:
// 		fmt.Fprintf(buf, "%q", v.String())
// 	case reflect.Bool:
// 		if v.Bool() {
// 			buf.WriteByte('t')
// 		} else {
// 			fmt.Fprint(buf, "nil")
// 		}
// 	case reflect.Ptr:
// 		return encode(indent, buf, v.Elem())
// 	case reflect.Interface: // ("[]int"(1 2 3))
// 		buf.WriteByte('(')
// 		fmt.Fprintf(buf, "%q", v.Elem().Type())
// 		if err := encode(indent, buf, v.Elem()); err != nil {
// 			return err
// 		}
// 		buf.WriteByte(')')
// 	case reflect.Array, reflect.Slice: // (value ...)
// 		buf.WriteByte('(')
// 		indent += 1
// 		for i := 0; i < v.Len(); i++ {
// 			if i > 0 {
// 				buf.WriteByte('\n')
// 				for i := 0; i < indent; i++ {
// 					buf.WriteByte(' ')
// 				}
// 			}
// 			if err := encode(indent, buf, v.Index(i)); err != nil {
// 				return err
// 			}
// 		}
// 		buf.WriteByte(')')
// 		indent -= 1
// 	case reflect.Struct: // ((name value) ...)
// 		buf.WriteByte('(')
// 		indent += 1
// 		for i := 0; i < v.NumField(); i++ {
// 			if i > 0 {
// 				buf.WriteByte('\n')
// 				for i := 0; i < indent; i++ {
// 					buf.WriteByte(' ')
// 				}
// 			}
// 			name := v.Type().Field(i).Name
// 			fmt.Fprintf(buf, "(%s ", name)
// 			indent += len(name) + 2
// 			if err := encode(indent, buf, v.Field(i)); err != nil {
// 				return err
// 			}
// 			indent -= len(name) + 2
// 			buf.WriteByte(')')
// 		}
// 		buf.WriteByte(')')
// 		indent -= 1
// 	case reflect.Map: // ((key value) ...)
// 		buf.WriteByte('(')
// 		indent += 1
// 		for i, key := range v.MapKeys() {
// 			if i > 0 {
// 				buf.WriteByte('\n')
// 				for i := 0; i < indent; i++ {
// 					buf.WriteByte(' ')
// 				}
// 			}
// 			buf.WriteByte('(')
// 			if err := encode(indent, buf, key); err != nil {
// 				return err
// 			}
// 			buf.WriteByte(' ')
// 			if err := encode(indent, buf, v.MapIndex(key)); err != nil {
// 				return err
// 			}
// 			buf.WriteByte(')')
// 		}
// 		buf.WriteByte(')')
// 		indent -= 1
// 	default: // float, compex, bool, chan, func, interface
// 		return fmt.Errorf("unsupported type: %s", v.Type())
// 	}
// 	return nil
// }
// func Marshal(v interface{}) ([]byte, error) {
// 	var buf bytes.Buffer
// 	var indent int
// 	if err := encode(indent, &buf, reflect.ValueOf(v)); err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }

func main() {
	sl, err := Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(sl))
}
