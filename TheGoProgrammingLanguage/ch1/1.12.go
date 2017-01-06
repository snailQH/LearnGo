//server2 is a minimal "echo" and counter server
package main

import(
	"fmt"
	"log"
	"net/http"
	"sync"
	"image"
	"image/gif"
	"io"
	"image/color"
	"math"
	"math/rand"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White,color.Black}

func main(){
	handler := func(w http.ResponseWriter, r *http.Request){
		lissajous(w)
	}
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))

}

//counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n",count)
	mu.Unlock()
}


const(
	whiteIndex = 0 //first color in palette
	redIndex = 1 //next color in palatte
)

func lissajous(out io.Writer){
	const(
		cycles = 20
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i< nframes; i++{
		rect := image.Rect(0,0,2*size+1,2*size+1)
		img := image.NewPaletted(rect,palette)
		for t := 0.0; t< cycles*2*math.Pi;t+=res{
			x := math.Sin(t)
			y := math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),redIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay,delay)
		anim.Image = append(anim.Image,img)
	}
	gif.EncodeAll(out,&anim) //NOTE: ingoring encoding errors
}
