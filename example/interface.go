package main

import "github.com/tikv/client-go/v2/tikv"

type Foo interface {
	Bar() tikv.Codec
}
