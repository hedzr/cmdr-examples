// Copyright © 2020 Hedzr Yeh.

package unsorted

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

// IsSlice test the type of 'v' if it is golang Slice.
func IsSlice(v interface{}) bool {
	var (
		// to      = s.indirect(reflect.ValueOf(toValue))
		to = reflect.ValueOf(v)
	)
	return to.Kind() == reflect.Slice
}

// CopySlice do a generic copy from slice to new slice.
func CopySlice(s interface{}) interface{} {
	t, v := reflect.TypeOf(s), reflect.ValueOf(s)
	c := reflect.MakeSlice(t, v.Len(), v.Len())
	reflect.Copy(c, v)
	return c.Interface()
}

// MergeSlice do a merging on generic slice type.
func MergeSlice(target interface{}, source interface{}) interface{} {
	v := reflect.ValueOf(source)
	tv := reflect.ValueOf(target)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}

	if tv.Kind() != reflect.Slice || v.Kind() != reflect.Slice {
		return target
	}

	if xm, ok := source.(map[interface{}]interface{}); ok {
		logrus.Debug("map", xm)
		return xm
	}

	for i := 0; i < v.Len(); i++ {
		value := v.Index(i)

		found := false
		for k := 0; k < tv.Len(); k++ {
			if tv.Index(k) == value {
				found = true
				break
			}
		}

		if !found {
			tv = reflect.Append(tv, value)
			reflect.ValueOf(&target).Elem().Set(tv)
		}
	}
	return target
}

// CloneMap do the clone on map[string]interface{}
func CloneMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		if vm, ok := v.(map[string]interface{}); ok {
			cp[k] = CloneMap(vm)
		} else {
			cp[k] = v
		}
	}
	return cp
}

// MergeMap do the recursive merge on a map[string]interface{}
func MergeMap(maps ...map[string]interface{}) map[string]interface{} {
	result, err := gormDefaultCopier.Copy(maps...)
	if err != nil {
		logrus.Error("merge map failed: ", err)
		return nil
	}
	return result
}

// Copier for depp-clone
type Copier interface {
	SetIgnoreNames(ignoreNames ...string) Copier
	Copy(maps ...map[string]interface{}) (result map[string]interface{}, err error)
}

type copierImpl struct {
	KeepIfFromIsNil  bool // 源字段值为nil指针时，目标字段的值保持不变
	ZeroIfEqualsFrom bool // 源和目标字段值相同时，目标字段被清除为未初始化的零值
	KeepIfFromIsZero bool // 源字段值为未初始化的零值时，目标字段的值保持不变 // 此条尚未实现
	IgnoreNames      []string
}

var (
	// merge two or more maps, not only deep-copy
	gormDefaultCopier = &copierImpl{KeepIfFromIsNil: true, ZeroIfEqualsFrom: true, KeepIfFromIsZero: true}
)

// SetIgnoreNames give a group of ignored fieldNames
func (s *copierImpl) SetIgnoreNames(ignoreNames ...string) Copier {
	s.IgnoreNames = ignoreNames
	return s
}

// Copy do a deep-clone on the giving 'maps'
func (s *copierImpl) Copy(maps ...map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	for _, from := range maps {
		for k, v := range from {
			if vm, ok := v.(map[string]interface{}); ok {
				toMap := make(map[string]interface{})
				toMap, err = s.Copy(toMap, vm)
				result[k] = toMap
			} else if IsSlice(v) {
				if result[k] == nil {
					result[k] = CopySlice(v)
				} else {
					// println("---- v.kind: ", reflect.ValueOf(v).Kind())
					result[k] = s.mergeSlice(result[k], v)
				}
			} else {
				if !s.KeepIfFromIsNil || v != nil {
					result[k] = v
				}
			}
		}
	}
	return
}

func (s *copierImpl) mergeGenericMap(to, from map[interface{}]interface{}) {
	for k, v := range from {
		if !s.KeepIfFromIsNil || v != nil {
			if vm, ok := v.(map[interface{}]interface{}); ok {
				if tvm, ok := to[k].(map[interface{}]interface{}); ok {
					s.mergeGenericMap(tvm, vm)
					continue
				}
			}
			to[k] = v
		}
	}
}

func (s *copierImpl) mergeSlice(target interface{}, source interface{}) interface{} {
	v := reflect.ValueOf(source)
	tv := reflect.ValueOf(target)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}

	if tv.Kind() != reflect.Slice || v.Kind() != reflect.Slice {
		return target
	}

	for i := 0; i < v.Len(); i++ {
		value := v.Index(i)
		// println("> value kind: ", value.Kind(), " | tv kind: ", tv.Kind())

		if vm, ok := value.Interface().(map[interface{}]interface{}); ok {
			// println(vm)
			// While the element i is a slice of map, trying to
			// merge the map into target[i] WITHOUT RECURSIVELY.
			if i < tv.Len() {
				if tvm, ok := tv.Index(i).Interface().(map[interface{}]interface{}); ok {
					s.mergeGenericMap(tvm, vm)
					continue
				}
			}
		}

		found := false
		for k := 0; k < tv.Len(); k++ {
			// println("  loop 'to' ", k, ": ", tv.Index(k).Interface(), " | kind = ", tv.Index(k).Kind())
			if reflect.DeepEqual(tv.Index(k).Interface(), value.Interface()) {
				found = true
				break
			}
		}

		if !found {
			tv = reflect.Append(tv, value)
			reflect.ValueOf(&target).Elem().Set(tv)
		}
	}
	return target
}
