package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	Title:    "", // "Dr. Strangelove",
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

// type    zero
// string  ""
// int 	0
// uint 	0
// bool    false
// slice
// struct
// map
// pointer
// interface{}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	// case reflect.Invalid:
	// 	fmt.Fprint(buf, "nil")
	case reflect.String:
		// buf.WriteString(v.String())
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	// case reflect.Float32, reflect.Float64:
	// case reflect.Complex128, reflect.Complex64:
	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "true")
		} else {
			fmt.Fprint(buf, "false")
		}
	case reflect.Slice, reflect.Array:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')
	case reflect.Struct:
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
	case reflect.Map:
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
	case reflect.Pointer:
		if v.IsNil() {
			fmt.Fprint(buf, "null")
		} else {
			if err := encode(buf, v.Elem()); err != nil {
				return err
			}
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprint(buf, "null")
		} else {
			if err := encode(buf, v.Elem()); err != nil {
				return err
			}
		}
	default:
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
func main() {
	s, _ := json.MarshalIndent(strangelove, "", " ")
	fmt.Println(string(s))
	// s, _ := Marshal(strangelove)
	// // fmt.Println(string(s))
	// sl := new(Movie)
	// err := json.Unmarshal(s, sl)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(sl)
}
