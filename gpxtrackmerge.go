package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/twpayne/go-geom"
	"github.com/twpayne/go-gpx"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		args = []string{
			"/Users/jum/Nextcloud/D & J/No Rush/Baja Winter 2022-2023.GPX",
		}
	}

	gpxFileArg := args[0]
	//gpxFile, err := gpx.ParseFile(gpxFileArg)
	f, err := os.Open(gpxFileArg)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g, err := gpx.Read(f)
	if err != nil {
		panic(err)
	}

	//spew.Dump(gpx)
	fmt.Printf("%#v\n", g)
	n := gpx.TrkType{Name: "Merged Tracks"}
	n.TrkSeg = []*gpx.TrkSegType{
		{},
	}
	for _, t := range g.Trk {
		fmt.Printf("%#v\n", t)
		for _, s := range t.TrkSeg {
			fmt.Printf("%#v\n", s)
			n.TrkSeg[0].TrkPt = append(n.TrkSeg[0].TrkPt, s.TrkPt...)
		}
	}
	// XXX sort by date?
	g.Trk = []*gpx.TrkType{&n}

	nf, err := os.Create("Merged.gpx")
	if err != nil {
		panic(err)
	}
	err = g.Write(nf)
	if err != nil {
		panic(err)
	}
	err = nf.Close()
	if err != nil {
		panic(err)
	}
}
