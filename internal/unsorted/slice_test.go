// Copyright © 2020 Hedzr Yeh.

package unsorted_test

import (
	"github.com/hedzr/cmdr-examples/internal/unsorted"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestMergeMap(t *testing.T) {
	var err error
	sTo := `
  test: [1,2,3]
  testArrMap:
    - x:
      c1: xc1
      z:
        s: s2
        t: 3zz 可以任意文字，任意类型，取决于你怎么使用 META 数据集
    - "y"  # cmdr will convert these value to boolean: 1,on,y,yes,t,true. so, quote it if you want a real string type value
`

	sFrom := `
test: [2,5]
testArrMap:
  - 
    c1:
    z:
      s: s3
    x: 1
  - "z"  # cmdr will convert these value to boolean: 1,on,y,yes,t,true. so, quote it if you want a real string type value
`
	from, to := make(map[string]interface{}), make(map[string]interface{})
	err = yaml.Unmarshal([]byte(sFrom), &from)
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(sTo), &to)
	if err != nil {
		t.Fatal(err)
	}

	to = unsorted.MergeMap(to, from)

	// // NOTE: merge CAN'T merge two map directly!!
	// if err := mergo.Merge(&to, from); err != nil {
	// 	t.Error("can't merge metadata: ", err)
	// }

	var b []byte
	b, err = yaml.Marshal(to)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))

	if len(to["test"].([]interface{})) != 4 {
		t.Fatalf("merge slice error, expect slice length is %v, but got %v", 4, len(to["test"].([]interface{})))
	}
}
