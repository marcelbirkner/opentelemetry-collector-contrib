// Copyright The OpenTelemetry Authors
// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println("starting pprof server")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	Execute()
}
