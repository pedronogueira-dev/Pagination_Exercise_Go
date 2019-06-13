package navigation

import (
	"testing"
)

type test struct {
	currentp   int
	totalp     int
	boundaries int
	around     int
	expected   string
}

func TestGetPagination(t *testing.T) {
	tests := []test{
		{1, 1, 0, 0, "1"},
		{2, 3, 0, 2, "1 2 3"},
		{4, 10, 2, 2, "1 2 3 4 5 6 ... 9 10"},
	}

	for _, test := range tests {
		res, err := getPagination(test.currentp, test.totalp, test.boundaries, test.around)
		if err != nil {
			t.Errorf("Error: %v", err)
			t.FailNow()
		}

		if test.expected != res {
			t.Errorf("Failed: \nexpected [%v]\nobtained[%v]", test.expected, res)
			t.FailNow()
		}
	}
}

func TestGetPagination_Boundaries(t *testing.T) {
	tests := []test{
		{1, 3, 1, 0, "1 ... 3"},
		{3, 5, 1, 0, "1 ... 3 ... 5"},
		{3, 5, 0, 0, "... 3 ..."},
	}

	for _, test := range tests {
		res, err := getPagination(test.currentp, test.totalp, test.boundaries, test.around)
		if err != nil {
			t.Errorf("Error: %v", err)
			t.FailNow()
		}

		if test.expected != res {
			t.Errorf("Failed: \nexpected [%v]\nobtained[%v]", test.expected, res)
			t.FailNow()
		}
	}
}

func TestGetPagination_Around(t *testing.T) {
	tests := []test{
		{1, 3, 0, 1, "1 2 ..."},
		{4, 4, 0, 1, "... 3 4"},
		{3, 5, 0, 1, "... 2 3 4 ..."},
	}

	for _, test := range tests {
		res, err := getPagination(test.currentp, test.totalp, test.boundaries, test.around)
		if err != nil {
			t.Errorf("Error: %v", err)
			t.FailNow()
		}

		if test.expected != res {
			t.Errorf("Failed: \nexpected [%v]\nobtained[%v]", test.expected, res)
			t.FailNow()
		}
	}
}

func TestGetPagination_CurrentPageInbound(t *testing.T) {
	tests := []test{
		{10, 5, 0, 0, ""},
		{0, 10, 0, 0, ""},
	}
	for _, test := range tests {
		_, err := getPagination(test.currentp, test.totalp, test.boundaries, test.around)
		if err != nil {
			t.Errorf("Error: %v", err)
			t.FailNow()
		}
	}
}
