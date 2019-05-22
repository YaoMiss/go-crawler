package parser

import (
	"go-crawler/engine"
	"regexp"
)

var profileRex = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRex = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	matches := profileRex.FindAllSubmatch(contents, -1)
	/**
	aTag[0] the a link
	aTag[1] the first ()
	aTag[2] the second ()
	*/
	for _, aTag := range matches {
		name := aTag[2]
		//result.Items = append(result.Items, "User "+string(name))
		result.Requests = append(result.Requests, engine.Request{Url: string(aTag[1]), ParserFunc: func(bytes []byte) engine.ParseResult {
			return ParseProfile(bytes, string(aTag[1]),  string(name))
		}})
	}

	matches = cityUrlRex.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
