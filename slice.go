/*
Slice相关工具
*/
package goutils

import (
	"errors"
	"reflect"
)

func Slice_DeleteByIndex(slicePtr interface{}, index int) error {

	slicePrtType := reflect.TypeOf(slicePtr)
	if slicePrtType.Kind() != reflect.Ptr {
		return errors.New("Error:Input param [slicePtr] must be pointer")
	}

	slicePtrValue := reflect.ValueOf(slicePtr)

	sliceValue := slicePtrValue.Elem()
	length := sliceValue.Len()
	if slicePtr == nil || length == 0 || (length-1) < index {
		sliceValue.Set(reflect.Zero(sliceValue.Type()))
	}
	if length-1 == index {
		//return sliceValue.Slice(0, index).Interface()
		sliceValue.Set(sliceValue.Slice(0, index))

	} else if (length - 1) >= index {
		//return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface()

		sliceValue.Set(reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)))
	}
	return nil
}
