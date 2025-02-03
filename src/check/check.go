package main

import (
	"fmt"

	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/base"
)

// This example shows how to
// 1. connect to a RTSP server
// 2. get and print informations about medias published on a path.

func main() {
	c := gortsplib.Client{}

	u, err := base.ParseURL("rtsp://888888:888888@103.59.133.26:554/asd22asdasd2asdsdawsd12e12e")
	if err != nil {
		panic(err)
	}

	err = c.Start(u.Scheme, u.Host)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	desc, _, err := c.Describe(u)

	if err != nil {
		panic(err)
	}
	for _, x := range desc.Medias {
		fmt.Println(x.Type, x.ID, x.Formats)

	}
}
