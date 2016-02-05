package vision

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type (
	// Param represents a valid parameter object
	Param struct {
		sync.RWMutex
		m map[string]interface{}
	}
	// Rules represents rules that applied to one single field
	Rules struct {
		sync.RWMutex
		m map[string]string
	}
)

// Set to set value to param
func (p *Param) Set(key string, val interface{}) {
	p.RLock()
	defer p.RUnlock()
	p.m[key] = val
}

// Get to retrieve value from param
func (p *Param) Get(key string) interface{} {
	p.RLock()
	defer p.RUnlock()
	return p.m[key]
}

// GetString to get string type value from param
func (p *Param) GetString(key string) string {
	val := p.Get(key)
	switch obj := val.(type) {
	case String:
		return obj.Val()
	case Integer:
		return obj.String()
	case Float:
		return obj.String()
	case Bool:
		return obj.String()
	case string:
		return obj
	case bool:
		return fmt.Sprintf("%t", obj)
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		return fmt.Sprintf("%d", obj)
	case float32, float64, complex64, complex128:
		return fmt.Sprintf("%f", obj)
	case nil:
		return ""
	default:
		panic(ErrUnsupported)
	}
	return ""
}

// GetInt to get int value from param
func (p *Param) GetInt(key string) int {
	val := p.Get(key)
	switch obj := val.(type) {
	case String:
		return obj.Integer().Val()
	case Integer:
		return obj.Val()
	case Float:
		return obj.Integer().Val()
	case Bool:
		return obj.Integer().Val()
	case string:
		i, err := strconv.Atoi(obj)
		if err != nil {
			panic(err)
		}
		return i
	case bool:
		if obj == true {
			return 1
		}
		return 0
	case int:
		return obj
	case int8:
		return int(obj)
	case int32:
		return int(obj)
	case int64:
		return int(obj)
	case uint:
		return int(obj)
	case uint8:
		return int(obj)
	case uint32:
		return int(obj)
	case uint64:
		return int(obj)
	case float32:
		return int(obj)
	case float64:
		return int(obj)
	case complex64, complex128:
		i, err := strconv.ParseInt(fmt.Sprintf("%v", obj), 10, 64)
		if err != nil {
			panic(err)
		}
		return int(i)
	case nil:
		return 0
	default:
		panic(ErrUnsupported)
	}
	return 0
}

// GetFloat to get float value from param
func (p *Param) GetFloat(key string) float64 {
	val := p.Get(key)
	switch obj := val.(type) {
	case String:
		return obj.Float().Val()
	case Integer:
		return obj.Float().Val()
	case Float:
		return obj.Val()
	case Bool:
		return obj.Float().Val()
	case string:
		f, err := strconv.ParseFloat(obj, 64)
		if err != nil {
			panic(err)
		}
		return f
	case bool:
		if obj == true {
			return 1.0
		}
		return 0.0
	case int:
		return float64(obj)
	case int8:
		return float64(obj)
	case int32:
		return float64(obj)
	case int64:
		return float64(obj)
	case uint:
		return float64(obj)
	case uint8:
		return float64(obj)
	case uint32:
		return float64(obj)
	case uint64:
		return float64(obj)
	case float32:
		return float64(obj)
	case float64:
		return obj
	case complex64, complex128:
		i, err := strconv.ParseFloat(fmt.Sprintf("%v", obj), 64)
		if err != nil {
			panic(err)
		}
		return float64(i)
	case nil:
		return 0.0
	default:
		panic(ErrUnsupported)
	}
	return 0.0
}

// GetBool to get bool value from param
func (p *Param) GetBool(key string) bool {
	val := p.Get(key)
	switch obj := val.(type) {
	case String:
		return obj.Bool().Val()
	case Integer:
		return obj.Bool().Val()
	case Float:
		return obj.Bool().Val()
	case Bool:
		return obj.Val()
	case string:
		str := strings.Trim(strings.ToLower(obj), Trim)
		if str == "true" {
			return true
		}
		return false
	case bool:
		return obj
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		switch obj {
		case 0:
			return false
		case 1:
			return true
		}
		return false
	case float32, float64, complex64:
		switch obj {
		case 1.0:
			return true
		case 0.0:
			return false
		}
		return false
	case nil:
		return false
	default:
		panic(ErrUnsupported)
	}
	return true
}
