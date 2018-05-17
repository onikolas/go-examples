package main

import "testing"

func TestDistance(t *testing.T) { type pair struct { a, b Point res float64 } tests := []pair{ {Point{1, 2}, Point{3, 4}, 2.82842712475}, {Point{3, 2}, Point{-3, -2}, 7.21110255093}, {Point{199, -999}, Point{0, 1}, 1019.60825811}, } for i := range tests { d := Distance(tests[i].a, tests[i].b) if !FloatEquals(d, tests[i].res, 0.0001) { t.Error(tests[i], d) } } }
