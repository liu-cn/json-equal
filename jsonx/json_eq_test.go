package jsonx

import (
	"fmt"
	"testing"
)

var tests = []test{
	{`{"uid":1,"nickname":"boyan","avatar":"avatar"}`, `{"nickname":"boyan","uid":1,"avatar":"avatar"}`, true},
	{`{"avatar":"avatar","nickname":"boyan","uid":1}`, `{"uid":1,"avatar":"avatar","nickname":"boyan"}`, true},
	{s1, s2, true},
	{s1, s2, true},
	{s1, s3, false},
	{s1, s4, false},
	{s1, s5, false},
	{s1, s6, false},
}

type test struct {
	json1 string
	json2 string
	want  bool
}

func ExampleEquals() {
	n1 := `{"LangAge":{"Name":"go"},"Name":"boyan"}`
	n2 := `{"Name":"boyan","LangAge":{"Name":"go"}}`
	n3 := `{"Name1":"boyan","Name":"boyan","LangAge":{"Name":"go"}}`
	fmt.Println(Equals(n1, n2, n3))
	fmt.Println(Equals(n1, n2))
	//Output false nil
	//Output true nil
}

func ExampleEqual() {
	j1 := `[{"id":1,"name":"c"},{"id":1,"name":"c++"}]`
	j2 := `[{"id":1,"name":"c"},{"id":1,"name":"c++"}]`
	j3 := `[{"id":1,"name":"c++"},{"id":1,"name":"c"}]`
	fmt.Println(Equal(j1, j2))
	fmt.Println(Equal(j2, j3))
	//Output true nil
	//Output false nil
}

func TestEq(t *testing.T) {
	for _, test := range tests {
		equal, err := Equals(test.json1, test.json2)
		if err != nil {
			t.Error(err)
		}
		if equal != test.want {
			t.Errorf("json eq err \njson1: %s\n json2:%s", test.json1, test.json2)
		}
	}
}
