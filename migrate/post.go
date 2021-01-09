package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type post struct {
	DateGmt      string   `xml:"date-gmt,attr"`
	Date         string   `xml:"date,attr"`
	Slug         string   `xml:"slug,attr"`
	RegularTitle string   `xml:"regular-title"`
	RegularBody  string   `xml:"regular-body"`
	Tag          []string `xml:"tag"`
}

type posts struct {
	Post []post `xml:"post"`
}

type tumblr struct {
	Posts posts `xml:"posts"`
}

func produceURL() string {
	params := map[string]string{
		"num":    "50",
		"type":   "text",
		"filter": "none",
	}
	return getURL("https://iamvolonbolon.tumblr.com/", "api/read/", params)
}

func getURL(base string, path string, params map[string]string) string {
	b, e := url.Parse(base)
	if e != nil {
		return ""
	}
	b.Path += path
	p := url.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	b.RawQuery = p.Encode()

	return b.String()
}

func getXMLFrom(url string) ([]byte, error) {
	r, e := http.Get(url)

	if e != nil {
		return []byte{}, fmt.Errorf("Error retrieving %v", url)
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", r.StatusCode)
	}

	d, pe := ioutil.ReadAll(r.Body)
	if pe != nil {
		return []byte{}, fmt.Errorf("Error processing response: %v", pe)
	}

	return d, nil
}

func (p post) formatHeader() string {
	str := "---\n"
	str += "layout: post\n"
	str += "title: \"" + p.RegularTitle + "\"\n"
	str += "date: " + p.DateGmt + "\n"
	str += "categories: " + strings.Join(p.Tag, " ") + "\n"
	str += "---\n"

	return str
}

func (p post) formatPost() string {
	header := p.formatHeader()
	body := p.RegularBody

	str := header + "\n" + body

	return str
}
