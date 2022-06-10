package main

import (
	//"fmt"
	"reflect"
	"testing"
)

func Test_number_Plus(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "Plus", fields: fields{12, 6}, want: 18},
		{name: "Plus", fields: fields{14, 7}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.Plus(); got != tt.want {
				t.Errorf("number.Plus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_Minus(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "Minus", fields: fields{12, 6}, want: 6},
		{name: "Minus", fields: fields{14, 7}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.Minus(); got != tt.want {
				t.Errorf("number.Minus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_Multi(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "Multi", fields: fields{12, 6}, want: 72},
		{name: "Multi", fields: fields{14, 7}, want: 98},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.Multi(); got != tt.want {
				t.Errorf("number.Multi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_Div(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{name: "Div", fields: fields{12, 6}, want: 1},
		{name: "Div", fields: fields{14, 7}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.Div(); got != tt.want {
				t.Errorf("number.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculator(t *testing.T) {
	n := number{12, 8}
	//var c int64
	calculator(n, 1)
	calculator(n, 2)
	calculator(n, 3)
	calculator(n, 4)
}

func Test_number_init(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
		want   number
	}{
		{name: "init", fields: fields{0, 0}, want: number{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			if got := n.init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("number.init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_number_show(t *testing.T) {
	type fields struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "show", fields: fields{14, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := number{
				num1: tt.fields.num1,
				num2: tt.fields.num2,
			}
			n.show()
		})
	}
}
