package main

import "fmt"

type RequestOptions struct {
	page    int
	perPage int
	sort    string
}

// Option型を引数にとる関数型 `Option` を定義します
type Option func(request *RequestOptions)

func Page(p int) Option {
	return func(r *RequestOptions) {
		if r != nil {
			r.page = p
		}
	}
}

func PerPage(pp int) Option {
	return func(r *RequestOptions) {
		if r != nil {
			r.perPage = pp
		}
	}
}

func Sort(s string) Option {
	return func(r *RequestOptions) {
		if r != nil {
			r.sort = s
		}
	}
}

func NewRequest(opts ...Option) *RequestOptions {
	// デフォルト値
	r := &RequestOptions{page: 1, perPage: 30, sort: "desc"}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func main() {
	r := NewRequest()
	// `Pageを設定`
	r = NewRequest(Page(10))
	fmt.Println(r)
	// `Page`と`PerPage``Sort`を設定`
	r = NewRequest(Page(10), PerPage(2), Sort("asc"))
	fmt.Println(r)
}
