package easygo

type Tuple1[T1 any] struct {
	V1 T1
}

func MakeTuple1[T1 any](v1 T1) Tuple1[T1] {
	return Tuple1[T1]{
		V1: v1,
	}
}

func NewTuple1[T1 any](v1 T1) *Tuple1[T1] {
	return &Tuple1[T1]{
		V1: v1,
	}
}

type Tuple2[T1 any, T2 any] struct {
	V1 T1
	V2 T2
}

func MakeTuple2[T1 any, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{
		V1: v1,
		V2: v2,
	}
}

func NewTuple2[T1 any, T2 any](v1 T1, v2 T2) *Tuple2[T1, T2] {
	return &Tuple2[T1, T2]{
		V1: v1,
		V2: v2,
	}
}

type Tuple3[T1 any, T2 any, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

func MakeTuple3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{
		V1: v1,
		V2: v2,
		V3: v3,
	}
}

func NewTuple3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3) *Tuple3[T1, T2, T3] {
	return &Tuple3[T1, T2, T3]{
		V1: v1,
		V2: v2,
		V3: v3,
	}
}

type Tuple4[T1 any, T2 any, T3 any, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

func MakeTuple4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
	}
}

func NewTuple4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) *Tuple4[T1, T2, T3, T4] {
	return &Tuple4[T1, T2, T3, T4]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
	}
}

type Tuple5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

func MakeTuple5[T1 any, T2 any, T3 any, T4 any, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
	}
}

func NewTuple5[T1 any, T2 any, T3 any, T4 any, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) *Tuple5[T1, T2, T3, T4, T5] {
	return &Tuple5[T1, T2, T3, T4, T5]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
	}
}

type Tuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

func MakeTuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Tuple6[T1, T2, T3, T4, T5, T6] {
	return Tuple6[T1, T2, T3, T4, T5, T6]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
	}
}

func NewTuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) *Tuple6[T1, T2, T3, T4, T5, T6] {
	return &Tuple6[T1, T2, T3, T4, T5, T6]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
	}
}

type Tuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
}

func MakeTuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return Tuple7[T1, T2, T3, T4, T5, T6, T7]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
	}
}

func NewTuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) *Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return &Tuple7[T1, T2, T3, T4, T5, T6, T7]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
	}
}

type Tuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
}

func MakeTuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
	}
}

func NewTuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return &Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
	}
}

type Tuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
}

func MakeTuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
	}
}

func NewTuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return &Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
	}
}

type Tuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
}

func MakeTuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
	}
}

func NewTuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return &Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
	}
}

type Tuple11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
}

func MakeTuple11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11) Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	return Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
	}
}

func NewTuple11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11) *Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	return &Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
	}
}

type Tuple12[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
}

func MakeTuple12[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12) Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12] {
	return Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
	}
}

func NewTuple12[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12) *Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12] {
	return &Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
	}
}

type Tuple13[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
}

func MakeTuple13[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13) Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13] {
	return Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
	}
}

func NewTuple13[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13) *Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13] {
	return &Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
	}
}

type Tuple14[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
}

func MakeTuple14[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14) Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14] {
	return Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
	}
}

func NewTuple14[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14) *Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14] {
	return &Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
	}
}

type Tuple15[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
}

func MakeTuple15[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15) Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15] {
	return Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
	}
}

func NewTuple15[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15) *Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15] {
	return &Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
	}
}

type Tuple16[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
}

func MakeTuple16[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16) Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16] {
	return Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
	}
}

func NewTuple16[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16) *Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16] {
	return &Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
	}
}

type Tuple17[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
}

func MakeTuple17[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17) Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17] {
	return Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
	}
}

func NewTuple17[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17) *Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17] {
	return &Tuple17[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
	}
}

type Tuple18[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
}

func MakeTuple18[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18) Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18] {
	return Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
	}
}

func NewTuple18[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18) *Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18] {
	return &Tuple18[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
	}
}

type Tuple19[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
}

func MakeTuple19[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19) Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19] {
	return Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
	}
}

func NewTuple19[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19) *Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19] {
	return &Tuple19[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
	}
}

type Tuple20[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
}

func MakeTuple20[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20) Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20] {
	return Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
	}
}

func NewTuple20[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20) *Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20] {
	return &Tuple20[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
	}
}

type Tuple21[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
}

func MakeTuple21[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21) Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21] {
	return Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
	}
}

func NewTuple21[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21) *Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21] {
	return &Tuple21[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
	}
}

type Tuple22[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
}

func MakeTuple22[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22) Tuple22[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22] {
	return Tuple22[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
	}
}

func NewTuple22[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22) *Tuple22[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22] {
	return &Tuple22[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
	}
}

type Tuple23[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
	V23 T23
}

func MakeTuple23[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23) Tuple23[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23] {
	return Tuple23[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
	}
}

func NewTuple23[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23) *Tuple23[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23] {
	return &Tuple23[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
	}
}

type Tuple24[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
	V23 T23
	V24 T24
}

func MakeTuple24[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24) Tuple24[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24] {
	return Tuple24[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
	}
}

func NewTuple24[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24) *Tuple24[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24] {
	return &Tuple24[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
	}
}

type Tuple25[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
	V23 T23
	V24 T24
	V25 T25
}

func MakeTuple25[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25) Tuple25[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25] {
	return Tuple25[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
	}
}

func NewTuple25[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25) *Tuple25[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25] {
	return &Tuple25[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
	}
}

type Tuple26[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
	V23 T23
	V24 T24
	V25 T25
	V26 T26
}

func MakeTuple26[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26) Tuple26[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26] {
	return Tuple26[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
	}
}

func NewTuple26[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26) *Tuple26[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26] {
	return &Tuple26[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
	}
}

type Tuple27[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
	V23 T23
	V24 T24
	V25 T25
	V26 T26
	V27 T27
}

func MakeTuple27[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26, v27 T27) Tuple27[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27] {
	return Tuple27[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
		V27: v27,
	}
}

func NewTuple27[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26, v27 T27) *Tuple27[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27] {
	return &Tuple27[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
		V27: v27,
	}
}

type Tuple28[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
	V23 T23
	V24 T24
	V25 T25
	V26 T26
	V27 T27
	V28 T28
}

func MakeTuple28[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26, v27 T27, v28 T28) Tuple28[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28] {
	return Tuple28[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
		V27: v27,
		V28: v28,
	}
}

func NewTuple28[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26, v27 T27, v28 T28) *Tuple28[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28] {
	return &Tuple28[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
		V27: v27,
		V28: v28,
	}
}

type Tuple29[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any, T29 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
	V23 T23
	V24 T24
	V25 T25
	V26 T26
	V27 T27
	V28 T28
	V29 T29
}

func MakeTuple29[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any, T29 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26, v27 T27, v28 T28, v29 T29) Tuple29[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28, T29] {
	return Tuple29[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28, T29]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
		V27: v27,
		V28: v28,
		V29: v29,
	}
}

func NewTuple29[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any, T29 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26, v27 T27, v28 T28, v29 T29) *Tuple29[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28, T29] {
	return &Tuple29[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28, T29]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
		V27: v27,
		V28: v28,
		V29: v29,
	}
}

type Tuple30[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any, T29 any, T30 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
	V10 T10
	V11 T11
	V12 T12
	V13 T13
	V14 T14
	V15 T15
	V16 T16
	V17 T17
	V18 T18
	V19 T19
	V20 T20
	V21 T21
	V22 T22
	V23 T23
	V24 T24
	V25 T25
	V26 T26
	V27 T27
	V28 T28
	V29 T29
	V30 T30
}

func MakeTuple30[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any, T29 any, T30 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26, v27 T27, v28 T28, v29 T29, v30 T30) Tuple30[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28, T29, T30] {
	return Tuple30[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28, T29, T30]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
		V27: v27,
		V28: v28,
		V29: v29,
		V30: v30,
	}
}

func NewTuple30[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any, T16 any, T17 any, T18 any, T19 any, T20 any, T21 any, T22 any, T23 any, T24 any, T25 any, T26 any, T27 any, T28 any, T29 any, T30 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16, v17 T17, v18 T18, v19 T19, v20 T20, v21 T21, v22 T22, v23 T23, v24 T24, v25 T25, v26 T26, v27 T27, v28 T28, v29 T29, v30 T30) *Tuple30[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28, T29, T30] {
	return &Tuple30[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18, T19, T20, T21, T22, T23, T24, T25, T26, T27, T28, T29, T30]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
		V6: v6,
		V7: v7,
		V8: v8,
		V9: v9,
		V10: v10,
		V11: v11,
		V12: v12,
		V13: v13,
		V14: v14,
		V15: v15,
		V16: v16,
		V17: v17,
		V18: v18,
		V19: v19,
		V20: v20,
		V21: v21,
		V22: v22,
		V23: v23,
		V24: v24,
		V25: v25,
		V26: v26,
		V27: v27,
		V28: v28,
		V29: v29,
		V30: v30,
	}
}

