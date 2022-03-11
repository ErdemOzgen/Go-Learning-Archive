package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
The default number generator is deterministic, so it’ll produce the same sequence
of numbers each time by default. To produce varying sequences, give it a seed that
changes. Note that this is not safe to use for random numbers you intend to be secret,
 use crypto/rand for those.
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	fmt.Println(x)

	var src cryptoSource
	rnd := rand.New(src)
	fmt.Println(rnd.Intn(1000)) // a truly random number 0 to 999 //https://yourbasic.org/golang/crypto-rand-int/

}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
