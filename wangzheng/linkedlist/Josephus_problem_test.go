package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNewCircleList(t *testing.T) {
// 	c := newCircleList(5)
// 	fmt.Println(c)
// }

func TestJosephusProblem(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{"test1", 0, 0, -1},
		{"test2", 0, 1, -1},
		{"test3", 0, 2, -1},
		{"test4", 1, 1, 1},
		{"test5", 1, 2, 1},
		{"test6", 1, 3, 1},
		{"test7", 2, 1, 2},
		{"test8", 2, 2, 1},
		{"test9", 2, 3, 2},
		{"test10", 3, 1, 3},
		{"test11", 3, 2, 3},
		{"test12", 3, 3, 2},
		{"test13", 40, 3, 28},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			survivor := JosephusProblem(tt.n, tt.k)
			assert.Equal(t, tt.want, survivor)
		})
	}
}

func TestJosephusProblem1(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
	}{
		{"t1", 1, 3},
		{"t2", 2, 3},
		{"t3", 3, 3},
		{"t4", 4, 3},
		{"t5", 5, 3},
		{"t6", 6, 3},
		{"t7", 7, 3},
		{"t8", 8, 3},
		{"t9", 9, 3},
		{"t10", 10, 3},
		{"t11", 11, 3},
		{"t12", 12, 3},
		{"t13", 13, 3},
		{"t14", 14, 3},
		{"t15", 15, 3},
		{"t16", 16, 3},
		{"t17", 17, 3},
		{"t18", 18, 3},
		{"t19", 19, 3},
		{"t20", 20, 3},
		{"t21", 21, 3},
		{"t22", 22, 3},
		{"t23", 23, 3},
		{"t24", 24, 3},
		{"t25", 25, 3},
		{"t26", 26, 3},
		{"t27", 27, 3},
		{"t28", 28, 3},
		{"t29", 29, 3},
		{"t30", 30, 3},
		{"t31", 31, 3},
		{"t32", 32, 3},
		{"t33", 33, 3},
		{"t34", 34, 3},
		{"t35", 35, 3},
		{"t36", 36, 3},
		{"t37", 37, 3},
		{"t38", 38, 3},
		{"t39", 39, 3},
		{"t40", 40, 3},
		{"t41", 41, 3},
		{"t42", 42, 3},
		{"t43", 43, 3},
		{"t44", 44, 3},
		{"t45", 45, 3},
		{"t46", 46, 3},
		{"t47", 47, 3},
		{"t48", 48, 3},
		{"t49", 49, 3},
		{"t50", 50, 3},
		{"t51", 51, 3},
		{"t52", 52, 3},
		{"t53", 53, 3},
		{"t54", 54, 3},
		{"t55", 55, 3},
		{"t56", 56, 3},
		{"t57", 57, 3},
		{"t58", 58, 3},
		{"t59", 59, 3},
		{"t60", 60, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			survivor := JosephusProblem(tt.n, tt.k)
			t.Logf("(%d, %d) get:%d", tt.n, tt.k, survivor)
		})
	}
}
