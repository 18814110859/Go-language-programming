package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r, err := os.Open("a")
	if err != nil {
		log.Fatal("error opening 'a'\n")
	}

	defer Close(r)

	a, err1 := os.Open("b")
	if err1 != nil {
		log.Fatal("error opening 'b'\n")
	}
	defer Close(a)
}

type Point struct {
	Longitude		string
	Latitude		string
	Distance		string
	ElevationGain	string
	ElevationLoss	string
}

func parse(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}
	
	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}

	return &p, nil
}

type Reader struct {
	r		io.Reader
	err 	error
}



func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func parseV2(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Latitude)
	r.read(&p.Longitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return &p, r.err
	}

	return &p, nil
}