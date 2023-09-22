// Package bignums provides utility functions for performing arithmetic operations on
// big numbers of various types including *big.Int and *big.Float.
package bignums

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strings"
)

type BigIntChain struct {
	value *big.Int
	stack []*big.Int
	err   error
}

func NewBigIntChain(val interface{}) *BigIntChain {
	bigInt, err := convertToBigInt(val)
	return &BigIntChain{value: bigInt, err: err}
}

func (bc *BigIntChain) Begin() *BigIntChain {
	bc.stack = append(bc.stack, new(big.Int).Set(bc.value))
	bc.value = big.NewInt(0)
	return bc
}

func (bc *BigIntChain) End() *BigIntChain {
	if len(bc.stack) == 0 {
		bc.err = fmt.Errorf("mismatched brackets")
		return bc
	}

	lastValue := bc.stack[len(bc.stack)-1]
	bc.stack = bc.stack[:len(bc.stack)-1]

	bc.value = new(big.Int).Add(lastValue, bc.value)
	return bc
}

func (bc *BigIntChain) Add(val interface{}) *BigIntChain {
	return bc.operate(val, func(a, b *big.Int) *big.Int { return new(big.Int).Add(a, b) })
}

func (bc *BigIntChain) Subtract(val interface{}) *BigIntChain {
	return bc.operate(val, func(a, b *big.Int) *big.Int { return new(big.Int).Sub(a, b) })
}

func (bc *BigIntChain) Multiply(val interface{}) *BigIntChain {
	return bc.operate(val, func(a, b *big.Int) *big.Int { return new(big.Int).Mul(a, b) })
}

func (bc *BigIntChain) Divide(val interface{}) *BigIntChain {
	return bc.operate(val, func(a, b *big.Int) *big.Int {
		if b.Cmp(big.NewInt(0)) == 0 {
			bc.err = fmt.Errorf("division by zero")
			return a // return the original value in case of an error
		}
		return new(big.Int).Div(a, b)
	})
}

func (bc *BigIntChain) Mod(val interface{}) *BigIntChain {
	return bc.operate(val, func(a, b *big.Int) *big.Int {
		if b.Cmp(big.NewInt(0)) == 0 {
			bc.err = fmt.Errorf("modulo by zero")
			return a
		}
		return new(big.Int).Mod(a, b)
	})
}

func (bc *BigIntChain) Pow(val interface{}) *BigIntChain {
	return bc.operate(val, func(a, b *big.Int) *big.Int {
		if uint64(b.Int64()) > uint64(64) {
			bc.err = fmt.Errorf("exponent too large")
			return a
		}
		return new(big.Int).Exp(a, b, nil)
	})
}

func (bc *BigIntChain) Abs() *BigIntChain {
	bc.value = new(big.Int).Abs(bc.value)
	return bc
}

func (bc *BigIntChain) operate(val interface{}, operation func(*big.Int, *big.Int) *big.Int) *BigIntChain {
	if bc.err != nil {
		return bc
	}
	bigInt, err := convertToBigInt(val)
	if err != nil {
		bc.err = err
		return bc
	}
	bc.value = operation(new(big.Int).Set(bc.value), bigInt)
	return bc
}

func (bc *BigIntChain) Value() (*big.Int, error) {
	if len(bc.stack) != 0 {
		return nil, fmt.Errorf("mismatched brackets")
	}

	return bc.value, bc.err
}

type BigFloatChain struct {
	value *big.Float
	stack []*big.Float
	err   error
}

func NewBigFloatChain(val interface{}) *BigFloatChain {
	bigFloat, err := convertToBigFloat(val)
	return &BigFloatChain{value: bigFloat, err: err}
}

func (bfc *BigFloatChain) Begin() *BigFloatChain {
	bfc.stack = append(bfc.stack, new(big.Float).Copy(bfc.value))
	bfc.value = big.NewFloat(0)
	return bfc
}

func (bfc *BigFloatChain) End() *BigFloatChain {
	if len(bfc.stack) == 0 {
		bfc.err = fmt.Errorf("mismatched brackets")
		return bfc
	}

	lastValue := bfc.stack[len(bfc.stack)-1]
	bfc.stack = bfc.stack[:len(bfc.stack)-1]

	bfc.value = new(big.Float).Add(lastValue, bfc.value)
	return bfc
}

func (bfc *BigFloatChain) Add(val interface{}) *BigFloatChain {
	return bfc.operate(val, func(a, b *big.Float) *big.Float { return new(big.Float).Add(a, b) })
}

func (bfc *BigFloatChain) Subtract(val interface{}) *BigFloatChain {
	return bfc.operate(val, func(a, b *big.Float) *big.Float { return new(big.Float).Sub(a, b) })
}

func (bfc *BigFloatChain) Multiply(val interface{}) *BigFloatChain {
	return bfc.operate(val, func(a, b *big.Float) *big.Float { return new(big.Float).Mul(a, b) })
}

func (bfc *BigFloatChain) Divide(val interface{}) *BigFloatChain {
	return bfc.operate(val, func(a, b *big.Float) *big.Float {
		if b.Cmp(big.NewFloat(0)) == 0 {
			bfc.err = fmt.Errorf("division by zero")
			return a // return the original value in case of an error
		}
		return new(big.Float).Quo(a, b)
	})
}

func (bfc *BigFloatChain) Pow(val interface{}) *BigFloatChain {
	return bfc.operate(val, func(a, b *big.Float) *big.Float {
		// Convert b to float64 for the exponent.
		exponent, _ := b.Float64()

		if exponent < 0 {
			bfc.err = fmt.Errorf("negative exponent")
			return a
		}

		if exponent != float64(int64(exponent)) {
			bfc.err = fmt.Errorf("non-integer exponent")
			return a
		}

		base, _ := a.Float64()

		// Use math.Pow to calculate power and then convert it back to *big.Float
		return new(big.Float).SetFloat64(math.Pow(base, exponent))
	})
}

func (bfc *BigFloatChain) Abs() *BigFloatChain {
	bfc.value = new(big.Float).Abs(bfc.value)
	return bfc
}

func (bfc *BigFloatChain) operate(val interface{}, operation func(*big.Float, *big.Float) *big.Float) *BigFloatChain {
	if bfc.err != nil {
		return bfc
	}
	bigFloat, err := convertToBigFloat(val)
	if err != nil {
		bfc.err = err
		return bfc
	}
	bfc.value = operation(new(big.Float).Copy(bfc.value), bigFloat)
	return bfc
}

func (bfc *BigFloatChain) Value() (*big.Float, error) {
	if len(bfc.stack) != 0 {
		return nil, fmt.Errorf("mismatched brackets")
	}

	return bfc.value, bfc.err
}

// convertToBigInt converts a supported type to a *big.Int.
// The function uses type assertion to determine the type of the input value.
// In case of *big.Float, the decimal part is truncated.
// For float32 and float64, the decimal part is truncated as well.
// In case of string, the function tries to determine whether the string is in decimal or hexadecimal form.
// If the string is hexadecimal, it is converted to decimal before conversion to *big.Int.
// If conversion is not possible, an error is returned.
func convertToBigInt(val interface{}) (*big.Int, error) {
	// Use type assertion to determine the type of the input value.
	switch v := val.(type) {
	case *big.Int:
		return v, nil
	case *big.Float:
		result := new(big.Int)
		v.Int(result)
		return result, nil
	case int, int64, int32, int16, int8:
		return big.NewInt(reflect.ValueOf(v).Int()), nil
	case uint64:
		return new(big.Int).SetUint64(v), nil
	case uint, uint32, uint16, uint8:
		return big.NewInt(int64(reflect.ValueOf(v).Uint())), nil
	case float32, float64:
		return big.NewInt(int64(reflect.ValueOf(v).Float())), nil
	case string:
		// If the string is hexadecimal, remove the prefix.
		base := 10
		if strings.HasPrefix(v, "0x") || strings.HasPrefix(v, "0X") {
			base = 16
			v = strings.TrimPrefix(strings.TrimPrefix(v, "0x"), "0X")
		}
		bigInt, success := new(big.Int).SetString(v, base)
		if !success {
			return nil, fmt.Errorf("could not convert string to big.Int")
		}
		return bigInt, nil
	default:
		return nil, fmt.Errorf("unsupported type: %v", reflect.TypeOf(val))
	}
}

// convertToBigFloat converts a supported type to a *big.Float.
func convertToBigFloat(val interface{}) (*big.Float, error) {
	// Use type assertion to determine the type of the input value.
	switch v := val.(type) {
	case *big.Float:
		return v, nil
	case *big.Int:
		return new(big.Float).SetInt(v), nil
	case int, int64, int32, int16, int8:
		return big.NewFloat(float64(reflect.ValueOf(v).Int())), nil
	case uint, uint64, uint32, uint16, uint8:
		return big.NewFloat(float64(reflect.ValueOf(v).Uint())), nil
	case float32, float64:
		return big.NewFloat(reflect.ValueOf(v).Float()), nil
	case string:
		// If the string is hexadecimal, convert it to big.Int first and then to big.Float.
		base := 10
		if strings.HasPrefix(v, "0x") || strings.HasPrefix(v, "0X") {
			base = 16
			v = strings.TrimPrefix(strings.TrimPrefix(v, "0x"), "0X")
		}
		if base == 16 {
			bigInt, success := new(big.Int).SetString(v, 16)
			if !success {
				return nil, fmt.Errorf("could not convert hex string to big.Float")
			}
			return new(big.Float).SetInt(bigInt), nil
		}
		bigFloat, success := new(big.Float).SetString(v)
		if !success {
			return nil, fmt.Errorf("could not convert string to big.Float")
		}
		return bigFloat, nil
	default:
		return nil, fmt.Errorf("unsupported type: %v", reflect.TypeOf(val))
	}
}
