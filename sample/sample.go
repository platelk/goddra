package main

import (
	"context"
	"fmt"
	"goddra/assetmanager"
	"goddra/assetmanager/fetcher"
	"goddra/geom"
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
	w, err := webglrender.NewWebGLRender(canvasEl)
	if err != nil {
		fmt.Println(err)
		return
	}
	rec := geom.NewRectangle(150, 400)
	rec.SetPosition(geom.Vec2{50, 50})
	fmt.Println(w.DrawRectangle(rec))
	fmt.Println(w.Render())
}

func main() {
	Run()
}