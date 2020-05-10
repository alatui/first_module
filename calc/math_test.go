package calc_test

import (
	"testing"

	"github.com/alatui/first_module/v2/calc"
)

func TestAdd(t *testing.T) {

	simpleSum, _ := calc.Add(1, 2)
	if simpleSum != 3 {
		t.Errorf("simpleSum failed, expected %v, got %v", 3, simpleSum)
	}

	manyArgumentsSum, _ := calc.Add(1, 2, 3)
	if manyArgumentsSum != 6 {
		t.Errorf("manyArgumentsSum failed, expected %v, got %v", 3, manyArgumentsSum)
	}

	_, err := calc.Add()
	if err == nil {
		t.Errorf("sumWithError failed, expected %v, got %v", "err", err)
	}

}
