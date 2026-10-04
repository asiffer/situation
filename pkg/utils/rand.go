package utils

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func fallBackRandomByte() byte {
	ch := make(chan byte)
	var b byte = 0
	go func() {
		for range 8 {
			ch <- 0xFF
		}
	}()

	for i := range 8 {
		select {
		case <-ch:
			// add 2^i
			b += 1 << i
		case <-ch:
			// nothing
		}
	}
	return b
}

func fallbackFillRandom(buffer []byte) {
	for i := range buffer {
		buffer[i] = fallBackRandomByte()
	}
}

func fillRandom(buffer []byte) {
	// Note that err == nil only if we read len(buffer) bytes.
	if _, err := rand.Read(buffer); err != nil {
		fallbackFillRandom(buffer)
	}
}

// RandUint16 returns an integer between 0 and max
func RandUint16(max uint16) uint16 {
	buffer := make([]byte, 2)
	fillRandom(buffer)
	return binary.BigEndian.Uint16(buffer) % max
}

func RandBytes(size int) []byte {
	buffer := make([]byte, size)
	fillRandom(buffer)
	return buffer
}

// func Rand128() [16]byte {
// 	var arr [16]byte
// 	buffer := make([]byte, 16)
// 	fillRandom(buffer)
// 	copy(arr[:], buffer)
// 	return arr
// }

// RandomTCPPort returns a TCP port between a and b
// a <= port < b
func RandomTCPPort(a, b uint16) uint16 {
	if a > b {
		a, b = b, a
	}
	// e := uint16(rand.Intn(int(b - a)))
	e := RandUint16(b - a)
	return a + e
}

var adjectives = []string{
	"admiring", "affectionate", "amazing", "awesome", "blissful",
	"bold", "brave", "brilliant", "charming", "cheerful",
	"clever", "compassionate", "confident", "cool", "curious",
	"dazzling", "determined", "dreamy", "eager", "ecstatic",
	"elegant", "energetic", "epic", "fearless", "focused",
	"friendly", "gallant", "gentle", "happy", "heuristic",
	"inspiring", "jolly", "keen", "lucid", "magnificent",
	"mystifying", "nifty", "optimistic", "peaceful", "quirky",
	"radiant", "relaxed", "serene", "sharp", "tremendous",
	"upbeat", "vibrant", "wizardly", "witty", "zealous",
}

var scientists = []string{
	"archimedes", "avogadro", "babbage", "bohr", "boltzmann",
	"copernicus", "curie", "darwin", "dijkstra", "dirac",
	"einstein", "erdos", "euler", "faraday", "fermi",
	"feynman", "franklin", "galileo", "gauss", "goodall",
	"hawking", "heisenberg", "hilbert", "hopper", "hubble",
	"hypatia", "johnson", "kepler", "knuth", "lamarr",
	"lavoisier", "leibniz", "lovelace", "maxwell", "meitner",
	"mendel", "mendeleev", "newton", "noether", "pasteur",
	"pauling", "planck", "ramanujan", "riemann", "rutherford",
	"schrodinger", "shannon", "tesla", "turing", "volta",
}

var _adjectivesCount = uint16(len(adjectives))

var _scientistsCount = uint16(len(scientists))

// RandomName returns a name like "tremendous erdos".
func RandomName() string {
	adj := adjectives[RandUint16(_adjectivesCount)]
	sci := scientists[RandUint16(_scientistsCount)]
	return fmt.Sprintf("%s %s", adj, sci)
}
