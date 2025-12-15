/*
Package slicesx provides generic slice utility functions.
*/
package slicesx

import (
	"cmp"
	"slices"
)

// Props returns a slice of a property, which is accessed using the given getter.
func Props[E, P any](vals []E, getter func(v E) P) []P {
	props := make([]P, len(vals))
	for i, v := range vals {
		props[i] = getter(v)
	}
	return props
}

// PropMap returns a map, mapping from a property, which is accessed using the given getter.
func PropMap[E any, P comparable](vals []E, getter func(v E) P) map[P]E {
	m := make(map[P]E, len(vals))
	for _, v := range vals {
		m[getter(v)] = v
	}
	return m
}

// PropsMap returns a map, mapping from a property, which is accessed using the given getter.
// It is allowed / normal that multiple elements have the same property value, so map values are slices collecting those elements.
func PropsMap[S ~[]E, E any, P comparable](vals S, getter func(v E) P) map[P]S {
	m := make(map[P]S)
	for _, v := range vals {
		p := getter(v)
		m[p] = append(m[p], v)
	}
	return m
}

// Filter returns a new slice holding only the filtered elements.
func Filter[S ~[]E, E any](vals S, f func(v E) bool) S {
	var out S
	for _, v := range vals {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// SortByPropValueOrder sorts a slice according to the order of property values specified in propOrder.
//
// If propOrder does not contain a property value found in the slice, the element will be moved to the end.
// If a property value is listed multiple times in propOrder, any of the indices may be used.
func SortByPropValueOrder[S ~[]E, E any, P comparable](vals S, propGetter func(v E) P, propOrder []P) {
	propIndex := make(map[P]int, len(propOrder))
	// Indexing propIndex with an unlisted value in propOrder results in the zero value (0),
	// shift indices by len(propOrder) to ensure missing props get higher index so they go to the end when sorting
	for i, p := range propOrder {
		propIndex[p] = i - len(propOrder)
	}

	slices.SortFunc(vals, func(v1, v2 E) int {
		return cmp.Compare(propIndex[propGetter(v1)], propIndex[propGetter(v2)])
	})
}

// Index "safely" indexes a slice.
// If the index is invalid (out of range) for the given slice, the (first) def is returned.
// If def is not specified, the zero value of E is returned.
func Index[E any](vals []E, idx int, def ...E) (result E) {
	if idx >= 0 && idx < len(vals) {
		return vals[idx]
	}
	if len(def) > 0 {
		return def[0]
	}
	return
}

// SelectByIndices selects elements from a slice specified by their indices.
// indices must hold valid indices, else a runtime panic may occur.
func SelectByIndices[S ~[]E, E any](vals S, indices []int) S {
	result := make(S, len(indices))

	for i, idx := range indices {
		result[i] = vals[idx]
	}

	return result
}
