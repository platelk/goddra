package main

import (
	"context"
	"fmt"
	"goddra/assetmanager"
	"goddra/assetmanager/fetcher"
	"goddra/render/webglrender"
	"syscall/js"
)

func Run() {
	f := fetcher.NewJsFetcher("http://localhost:8080")
	a := assetmanager.AssetManager{}
	a.AddFetcher(0, f)

	fmt.Println("Fetching ...")
	as, err := a.Fetch(context.Background(), "index.html")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(as.DataString())
	}
	// Init Canvas stuff
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "gocanvas")
	w, err := webglrender.NewWebGL(canvasEl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(w.DrawTriangle())
}

func main() {
	Run()
}