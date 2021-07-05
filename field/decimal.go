package field

import (
	"github.com/shopspring/decimal"
	"reflect"
)

type Decimal struct {
	decimal.NullDecimal
}

var TypeDecimal = reflect.TypeOf(Decimal{})