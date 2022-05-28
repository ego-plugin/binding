package fields

import (
	"bytes"
	"github.com/shopspring/decimal"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
)

type Decimal struct {
	decimal.NullDecimal
}

func (d *Decimal) NilValue() *Decimal {
	if !d.Valid {
		return nil
	}
	return d
}

func (d *Decimal) Ptr() *decimal.NullDecimal {
	if !d.Valid {
		return nil
	}
	return &d.NullDecimal
}

// UnmarshalMsgpack implements the msgpack.UnmarshalMsgpack interface.
func (d *Decimal) UnmarshalMsgpack(b []byte) error {
	if bytes.Equal(b, nullString) {
		return d.Scan(nil)
	}
	var s any
	if err := msgpack.Unmarshal(b, &s); err != nil {
		d.Valid = false
		return err
	}
	return d.Scan(s)
}

// MarshalMsgpack implements the msgpack.MarshalMsgpack interface.
func (d Decimal) MarshalMsgpack() ([]byte, error) {
	if !d.Valid {
		return nullString, nil
	}
	return d.MarshalText()
}

var (
	_           ValueScanner = (*Decimal)(nil)
	DecimalType              = reflect.TypeOf(Decimal{})
)
