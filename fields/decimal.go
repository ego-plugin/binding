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

// UnmarshalMsgpack implements the msgpack.UnmarshalMsgpack interface.
func (d *Decimal) UnmarshalMsgpack(b []byte) error {
	// scan for null
	if bytes.Equal(b, nullString) {
		return d.Scan(nil)
	}
	var s interface{}
	if err := msgpack.Unmarshal(b, &s); err != nil {
		return err
	}
	return d.Scan(s)
}

// MarshalMsgpack implements the msgpack.MarshalMsgpack interface.
func (d Decimal) MarshalMsgpack() ([]byte, error) {
	if !d.Valid {
		return nullString, nil
	}
	return msgpack.Marshal(d.Decimal.String())
}

var (
	_           ValueScanner = (*Decimal)(nil)
	DecimalType              = reflect.TypeOf(Decimal{})
)
