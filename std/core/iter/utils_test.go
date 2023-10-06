// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package iter

import "testing"

func TestIndex(t *testing.T) {
	type args struct {
		i int
		n int
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
		wantValid bool
	}{
		{
			name:      "index > sz",
			args:      args{1, 0},
			wantIndex: 1,
			wantValid: false,
		},
		{
			name:      "index = sz",
			args:      args{1, 1},
			wantIndex: 1,
			wantValid: false,
		},
		{
			name:      "neg index > sz",
			args:      args{-2, 0},
			wantIndex: -1,
			wantValid: false,
		},
		{
			name:      "neg index > sz",
			args:      args{-3, 1},
			wantIndex: -2,
			wantValid: false,
		},
		{
			name:      "index < sz",
			args:      args{1, 2},
			wantIndex: 1,
			wantValid: true,
		},
		{
			name:      "index < sz",
			args:      args{0, 1},
			wantIndex: 0,
			wantValid: true,
		},
		{
			name:      "neg index < sz",
			args:      args{-2, 3},
			wantIndex: 1,
			wantValid: true,
		},
		{
			name:      "neg index < sz",
			args:      args{-3, 3},
			wantIndex: 0,
			wantValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotValid := Index(tt.args.i, tt.args.n)
			if gotIndex != tt.wantIndex {
				t.Errorf("Index() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotValid != tt.wantValid {
				t.Errorf("Index() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func TestBound(t *testing.T) {
	type args struct {
		i int
		n int
	}
	tests := []struct {
		name      string
		args      args
		wantBound int
		wantValid bool
	}{
		{
			name:      "bound > sz",
			args:      args{1, 0},
			wantBound: 1,
			wantValid: false,
		},
		{
			name:      "bound = sz",
			args:      args{1, 1},
			wantBound: 1,
			wantValid: true,
		},
		{
			name:      "neg bound > sz",
			args:      args{-2, 0},
			wantBound: -1,
			wantValid: false,
		},
		{
			name:      "neg bound > sz",
			args:      args{-4, 1},
			wantBound: -2,
			wantValid: false,
		},
		{
			name:      "bound < sz",
			args:      args{1, 2},
			wantBound: 1,
			wantValid: true,
		},
		{
			name:      "bound < sz",
			args:      args{0, 1},
			wantBound: 0,
			wantValid: true,
		},
		{
			name:      "neg bound < sz",
			args:      args{-2, 3},
			wantBound: 2,
			wantValid: true,
		},
		{
			name:      "neg bound < sz",
			args:      args{-4, 3},
			wantBound: 0,
			wantValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBound, gotValid := Bound(tt.args.i, tt.args.n)
			if gotBound != tt.wantBound {
				t.Errorf("Bound() gotBound = %v, want %v", gotBound, tt.wantBound)
			}
			if gotValid != tt.wantValid {
				t.Errorf("Bound() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func TestStartEnd(t *testing.T) {
	type args struct {
		i int
		j int
		n int
	}
	tests := []struct {
		name        string
		args        args
		wantActuali int
		wantActualj int
		wantValid   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotActuali, gotActualj, gotValid := Range(tt.args.i, tt.args.j, tt.args.n)
			if gotActuali != tt.wantActuali {
				t.Errorf("StartEnd() gotActuali = %v, want %v", gotActuali, tt.wantActuali)
			}
			if gotActualj != tt.wantActualj {
				t.Errorf("StartEnd() gotActualj = %v, want %v", gotActualj, tt.wantActualj)
			}
			if gotValid != tt.wantValid {
				t.Errorf("StartEnd() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func TestEach(t *testing.T) {
	type args struct {
		iter SliceIter[int]
		fn   func(i int, v int) bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Each(tt.args.iter, tt.args.fn)
		})
	}
}

func TestEachEx(t *testing.T) {
	type args struct {
		iter Core[int]
		arg  float32
		fn   func(i int, v int, arg float32) bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EachEx(tt.args.iter, tt.args.arg, tt.args.fn)
		})
	}
}
