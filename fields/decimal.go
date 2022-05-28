package fields

import (
	"github.com/shopspring/decimal"
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
	return d.UnmarshalJSON(b)
}

// MarshalMsgpack implements the msgpack.MarshalMsgpack interface.
func (d Decimal) MarshalMsgpack() ([]byte, error) {
	return d.MarshalJSON()
}

var (
	_           ValueScanner = (*Decimal)(nil)
	DecimalType              = reflect.TypeOf(Decimal{})
)
