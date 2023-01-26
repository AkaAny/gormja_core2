package mirror

import (
	"unicode"
	_ "unsafe"
)

// name is an encoded type name with optional extra data.
//
// The first byte is a bit field containing:
//
//	1<<0 the name is exported
//	1<<1 tag data follows the name
//	1<<2 pkgPath nameOff follows the name and tag
//	1<<3 the name is of an embedded (a.k.a. anonymous) field
//
// Following that, there is a varint-encoded length of the name,
// followed by the name itself.
//
// If tag data is present, it also has a varint-encoded length
// followed by the tag itself.
//
// If the import path follows, then 4 bytes at the end of
// the data form a nameOff. The import path is only set for concrete
// methods that are defined in a different package than their type.
//
// If a name starts with "*", then the exported bit represents
// whether the pointed to type is exported.
//
// Note: this encoding must match here and in:
//   cmd/compile/internal/reflectdata/reflect.go
//   runtime/type.go
//   internal/reflectlite/type.go
//   cmd/link/internal/ld/decodesym.go

type name struct {
	bytes *byte
}

//go:linkname newName reflect.newName
func newName(n, tag string, exported, embedded bool) name

type nameOff int32 // offset to a name

//go:linkname resolveReflectName reflect.resolveReflectName
func resolveReflectName(n name) nameOff

func newNameOff(nameStr string) nameOff {
	var isExported = unicode.IsUpper(rune(nameStr[0]))
	var rrn = resolveReflectName(newName(nameStr, "", isExported, false))
	return rrn
}
