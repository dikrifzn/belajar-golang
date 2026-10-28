package helper

import (
"testing"
"github.com/stretchr/testify/assert"
_"fmt"
)

func TestHelloWorldAssertion(t *testing.T){
	result := HelloWorld("Dinda")
	assert.Equal(t,"Hello Dinda",result, "Result is not Hello Dinda")
	t.Run("sub", func(t *testing.T){
		result := HelloWorld("sub")
		assert.Equal(t,"Hello sub",result, "Result is not Hello sub")
	})
	t.Run("subF", func(t *testing.T){
		result := HelloWorld("subFail")
		assert.Equal(t,"Hello subF",result, "Result is not Hello subF")
	})
}

func TestHelloWorldDikri(t *testing.T){
	result := HelloWorld("Dikri")
	if result != "Hello Dikri"{
		t.Error("Result is not Hello Dikri")
	}
}

func TestHelloWorldDinda(t *testing.T){
	result := HelloWorld("Dinda")
	if result != "Hello Dinda"{
		t.Fatal("Result is not Hello Dinda")
	}
}

func TestTable(t *testing.T){
	testTable := []struct {
	name string
	request string
	expected string
}{
	{
		name: "dikri",
		request: "dikri",
		expected: "Hello dikri",
		
	},
	{
		name: "dinda",
		request: "dindaF",
		expected: "Hello dinda",
		
	},
}

for _, test := range testTable{
	t.Run(test.name, func(t *testing.T){
		result := HelloWorld(test.request)
		assert.Equal(t, test.expected,result)
	})
}
}