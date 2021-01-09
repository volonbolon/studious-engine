package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	xmlBytes, e := getXMLFrom(produceURL())
	if e != nil {
		os.Exit(1)
	}

	var d tumblr
	xml.Unmarshal(xmlBytes, &d)
	for i := 0; i < len(d.Posts.Post); i++ {
		post := d.Posts.Post[i]

		date := strings.Split(post.DateGmt, " ")[0]
		fn := date + "-" + post.Slug + ".markdown"

		p := post.formatPost()
		err := ioutil.WriteFile(fn, []byte(p), 0644)
		if err != nil {
			os.Exit(2)
		}
	}
}
