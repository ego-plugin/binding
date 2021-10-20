package field

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"reflect"
)

type Decimal struct {
	decimal.NullDecimal
}

var (
	_ json.Marshaler = (*Decimal)(nil)
	_ json.Unmarshaler = (*Decimal)(nil)
	TypeDecimal = reflect.TypeOf(Decimal{})
)