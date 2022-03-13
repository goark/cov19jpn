//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/goark/cov19jpn/chart"
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
	//prefcode := prefcodejpn.SHIMANE
	prefcode := prefcodejpn.TOKYO
	es, err := fetch.Import(
		r,
		filter.New(prefcode),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	list := entity.NewList(es)
	hlist := chart.New(list.StartDayMeasure(), list.EndDayMeasure().AddDay(7), 7, list)
	if err := chart.MakeHistChart(hlist, prefcode.Title(), "./chart-"+prefcode.Name()+".png"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
