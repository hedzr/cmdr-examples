/*
 */

// Copyright © 2020 Hedzr Yeh.

package tool

// PollardArray converts 'in' to new return object A.
func PollardArray(in []interface{}) interface{} {
	out := []interface{}{}
	for _, i := range in {
		out = append(out, Pollard(i))
	}
	return out
}

// Pollard 转换一个泛型对象中的全部 map[interface{}]interface{} 子对象为 map[string]interface{}
// 在转换完毕之后，新的返回对象将能够正确地被json编码：
// text = `age: 12
// name: joe`
// obj, err := yaml.Unmarshal(text)
// b, err := json.Marshal(Pollard(obj))
//
// 典型的用途在于将 golang 对象直接json输出时
func Pollard(in interface{}) interface{} {
	// var out interface{}
	if m, ok := in.(map[interface{}]interface{}); ok {
		o := make(map[string]interface{})
		for key, value := range m {
			switch k := key.(type) {
			case string:
				if a, ok := value.([]interface{}); ok {
					value = PollardArray(a)
				} else if mm, ok := value.(map[interface{}]interface{}); ok {
					value = Pollard(mm)
				}
				o[k] = value
			default:
				// o[key] = v
			}
		}
		return o
	}
	return in
}
