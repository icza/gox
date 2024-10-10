package gox

// Struct2 is a struct of 2 generic fields.
// May come handy as a "quick" wrapper for key fields or multiple results for [OpCache].
type Struct2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

func Struct2Of[T1, T2 any](v1 T1, v2 T2) Struct2[T1, T2] {
	return Struct2[T1, T2]{v1, v2}
}

// Struct3 is a struct of 3 generic fields.
// May come handy as a "quick" wrapper for key fields or multiple results for [OpCache].
type Struct3[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

func Struct3Of[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Struct3[T1, T2, T3] {
	return Struct3[T1, T2, T3]{v1, v2, v3}
}

// Struct4 is a struct of 4 generic fields.
// May come handy as a "quick" wrapper for key fields or multiple results for [OpCache].
type Struct4[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

func Struct4Of[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Struct4[T1, T2, T3, T4] {
	return Struct4[T1, T2, T3, T4]{v1, v2, v3, v4}
}

// Struct5 is a struct of 5 generic fields.
// May come handy as a "quick" wrapper for key fields or multiple results for [OpCache].
type Struct5[T1, T2, T3, T4, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

func Struct5Of[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Struct5[T1, T2, T3, T4, T5] {
	return Struct5[T1, T2, T3, T4, T5]{v1, v2, v3, v4, v5}
}

// Struct6 is a struct of 6 generic fields.
// May come handy as a "quick" wrapper for key fields or multiple results for [OpCache].
type Struct6[T1, T2, T3, T4, T5, T6 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

func Struct6Of[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Struct6[T1, T2, T3, T4, T5, T6] {
	return Struct6[T1, T2, T3, T4, T5, T6]{v1, v2, v3, v4, v5, v6}
}
