package database

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type entity interface {
	Content
}

var ignoreFieldNames = []string{"CreatedAt", "UpdatedAt"}

func EqualEntity[E entity](x, y E) bool {
	cmpoptIgnore := cmpopts.IgnoreFields(E{}, ignoreFieldNames...)
	return cmp.Equal(x, y, cmpoptIgnore)
}
func DiffEntity[E entity](x, y E) string {
	cmpoptIgnore := cmpopts.IgnoreFields(E{}, ignoreFieldNames...)
	return cmp.Diff(x, y, cmpoptIgnore)
}

func EqualEntities[E entity](x, y []E) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		w := y[i]
		if !EqualEntity(v, w) {
			return false
		}
	}
	return true
}
func DiffEntities[E entity](x, y []E) string {
	if len(x) != len(y) {
		return fmt.Sprintf("different size: %v, %v", len(x), len(y))
	}
	var result = ""
	for i, v := range x {
		w := y[i]
		result += DiffEntity(v, w)
	}
	return result
}
