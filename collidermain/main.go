// Copyright (c) 2014 The WebRTC project authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"collider/collider"
	"flag"
	"log"
)

var tls = flag.Bool("tls", false, "whether TLS is used")
var port = flag.Int("port", 443, "The TCP port that the server listens on")

func main() {
	flag.Parse()

	if *tls {
		log.Printf("Starting collider: https://localhost:%d", *port)
	} else {
		log.Printf("Starting collider: http://localhost:%d", *port)
	}
	c := collider.NewCollider()
	c.Run(*port, *tls)
}
