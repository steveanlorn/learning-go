package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ==================================================
// CORE PROBLEM API

// Data is the structure of the data we are copying.
type Data struct {
	Row string
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

// Hamm is a system we need to pull data from.
type Hamm struct {
	Host    string
	Timeout time.Duration
}

// Pull pulls data from Hamm.
func (h *Hamm) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 9, 10:
		return io.EOF
	case 5:
		return fmt.Errorf("error reading data")
	default:
		d.Row = "data"
		fmt.Println("In:", d.Row)
		return nil
	}
}

// Piglet is a system we need to store data into.
type Piglet struct {
	Host    string
	Timeout time.Duration
}

// Store stores data to Piglet.
func (p *Piglet) Store(d *Data) error {
	fmt.Println("Out:", d.Row)
	return nil
}

// ==================================================
// LOW-LEVEL API
// Composition of two behaviors of pull & store from the CORE API

// A non practical solution
// Because too tedious.

//type System2 struct {
//	Porky
//	Piglet
//}
//
//type System3 struct {
//	Porky
//	Pumbaa
//}

// pull knows how to pull bulks of data from Puller.
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		err := p.Pull(&data[i])
		if err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// store knows how to store bulks of data from Storer
func store(s Storer, data []Data) (int, error) {
	for i := range data {
		err := s.Store(&data[i])
		if err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// ==================================================
// HIGH-LEVEL API
// Do pulling and storing in a large batch

// Copy knows how to pull & store data from the System.
func Copy(p Puller, s Storer, batch int) error {
	data := make([]Data, batch) // <--- only one allocation throughout the code stack

	for {
		i, err := pull(p, data)
		if i > 0 {
			_, err := store(s, data)
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	puller := &Hamm{
		Host:    "localhost:8000",
		Timeout: time.Second,
	}

	storer := &Piglet{
		Host:    "localhost:9000",
		Timeout: time.Second,
	}

	if err := Copy(puller, storer, 3); err != io.EOF {
		fmt.Println(err)
	}
}
