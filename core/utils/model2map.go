package utils

import "reflect"

// 模型转map
func Model2Map(src any) map[string]any {
	ret := map[string]any{}
	tv := reflect.TypeOf(src)
	vv := reflect.ValueOf(src)
	l := tv.NumField()
	for i := 0; i < l; i++ {
		f := tv.Field(i)
		kname := f.Tag.Get("json")
		ret[kname] = vv.Field(i).Interface()
	}
	return ret
}

// 模型数组转map
func ModelList2Map(src any) []map[string]any {
	vv := reflect.ValueOf(src)

	ret := []map[string]any{}
	l := vv.Len()
	for i := 0; i < l; i++ {
		ret = append(ret, Model2Map(vv.Index(i).Interface()))
	}
	// for _, item := range src {
	// 	ret = append(ret, Model2Map(item))
	// }
	return ret
}
