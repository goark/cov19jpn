//go:build run
// +build run

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/goark/cov19jpn/entity"
	"github.com/goark/cov19jpn/fetch"
	"github.com/goark/cov19jpn/filter"
	"github.com/goark/cov19jpn/values/prefcodejpn"
)

func main() {
	r, err := fetch.Web(context.Background(), &http.Client{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer r.Close()
	es, err := fetch.Import(
		r,
		filter.New(prefcodejpn.SHIMANE),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	list := entity.NewList(es)
	list.Sort()
	_, _ = io.Copy(os.Stdout, bytes.NewReader(list.EncodeCSV()))
}
