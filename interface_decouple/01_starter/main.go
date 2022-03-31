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

// Start with concrete problem first not interface
// Interface should be discovered with refactoring

// ==================================================
// CORE PROBLEM API

// This exercise doesn't care about data transformation
// therefore we use a simple data structure representing a data.

// Focus on concrete data first not interface.

// Data is the structure of the data we are copying.
type Data struct {
	Row string
}

// Hamm is a system we need to pull data from.
type Hamm struct {
	Host    string
	Timeout time.Duration
}

// We use method here because Hamm is a stateful application.

// Drafting the code
// Does not mean that it the best code yet, but it works and readable.
// If it is not perform good enough, then we can improve it later with a help form tooling.

// Pull does not have memory allocation because potentially running in a tight loop.

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

// Drafting the code
// Does not mean choosing a system composition is a good idea now,
// but it is something come out in the mind now.

// System wraps Hamm & Piglet together into a single system
type System struct {
	Hamm
	Piglet
}

// function based API used here
// use function until it does not make sense
// We use function because system is not a state, it just cares about the behavior.

// pull knows how to pull bulks of data from Hamm.
func pull(h *Hamm, data []Data) (int, error) {
	for i := range data {
		err := h.Pull(&data[i])
		if err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// store knows how to store bulks of data from Piglet
func store(p *Piglet, data []Data) (int, error) {
	for i := range data {
		err := p.Store(&data[i])
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
func Copy(sys *System, batch int) error {
	data := make([]Data, batch) // <--- only one allocation throughout the code stack

	for {
		i, err := pull(&sys.Hamm, data)
		if i > 0 {
			_, err := store(&sys.Piglet, data)
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
	sys := System{
		Hamm{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Piglet{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
