//go:build go1.18 || go1.19
// +build go1.18 go1.19

package maphash

import (
	"reflect"
	"unsafe"
)

func getRuntimeHasher2[K any]() (h hashfn) {
	var k K
	a := reflect.MakeMap(reflect.MapOf(reflect.TypeOf(k), reflect.TypeOf(struct{}{}))).Interface()
	i := (*mapiface)(unsafe.Pointer(&a))
	h = i.typ.hasher
	return
}
