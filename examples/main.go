// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"net/http"

	query "github.com/dandycheung/html-query"
	. "github.com/dandycheung/html-query/expr"
)

func main() {
	etherscan()
}

func etherscan() {
	r := get(`https://cn.etherscan.com/address/0xb3dfb4ced9772f2c012265e755744deebb2891de`)
	defer r.Close()
	root, err := query.Parse(r)
	checkError(err)
	textCode := root.Pre(Id("editor")).Text()
	if textCode != nil {
		pn(*textCode)
	}
}

func golang() {
	r := get(`http://blog.golang.org/index`)
	defer r.Close()
	root, err := query.Parse(r)
	checkError(err)
	root.Div(Id("content")).Children(Class("blogtitle")).For(func(item *query.Node) {
		href := item.Ahref().Href()
		date := item.Span(Class("date")).Text()
		tags := item.Span(Class("tags")).Text()
		if href != nil {
			pn(*href)
		}
		if date != nil {
			pn(*date)
		}
		if tags != nil {
			p(*tags)
		}
	})
}

func get(url string) io.ReadCloser {
	resp, err := http.Get(url)
	checkError(err)
	return resp.Body
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func pn(v ...interface{}) {
	fmt.Print(v...)
}

func p(v ...interface{}) {
	fmt.Println(v...)
}
