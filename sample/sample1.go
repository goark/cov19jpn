//go:build run
// +build run

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/goark/cov19jpn/entity"
	"github.com/goark/cov19jpn/fetch"
	"github.com/goark/cov19jpn/filter"
	"github.com/goark/cov19jpn/values/date"
	"github.com/goark/cov19jpn/values/prefcodejpn"
)

func main() {
	r, err := fetch.File("./forecast_JAPAN_PREFECTURE_28.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer r.Close()
	es, err := fetch.Import(
		r,
		filter.New(prefcodejpn.SHIMANE).SetPeriod(date.NewPeriod(date.Zero, date.FromString("2021-01-24").AddDay(7+1))),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	list := entity.NewList(es)
	list.Sort()
	_, _ = io.Copy(os.Stdout, bytes.NewReader(list.EncodeCSV()))
}
