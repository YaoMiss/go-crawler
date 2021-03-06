package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const citySize = 470

	expectUrl := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(result.Requests) != citySize {
		t.Errorf("result should have %d "+"request,but had %d", citySize, len(result.Requests))
	}

	for i, url := range expectUrl {
		if result.Requests[i].Url != url {
			t.Errorf("expected url %d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
}
