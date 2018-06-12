package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/big"
	"time"
)

type Respuesta struct {
	login   string
	tranKey string
	seed    string
	nonce   string
}

var login = "usuarioprueba"
var trankey = "ABCD1234"
var _seed string
var _nonce string

func generate() {

	//nonce generado con base en: https://stackoverflow.com/questions/45428126/how-to-create-a-big-int-with-a-secure-random
	//Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
	}
	//String representation of n in base 32
	_nonce = n.Text(32)
	////////////////////////////////////////////

	now := time.Now()
	// Display the time as RFC3339
	_seed = now.Format(time.RFC3339)
}

func getRealNonce() string {
	return _nonce
}

func getLogin() string {
	return login
}

func getNonce() string {
	data2 := []byte(_nonce)
	dec := base64.StdEncoding.EncodeToString(data2)
	return dec
}

func getSeed() string {
	return _seed
}

func getTranKey() string {
	data := []byte(_nonce + _seed + trankey)
	hasher := sha1.New()
	hasher.Write(data)
	sha1_hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha1_hash
}

func setSeed(seed string) {
	_seed = seed
	fmt.Println(_seed)
}

func setNonce(nonce string) {
	_nonce = nonce
	fmt.Println(_nonce)
}

func asObject() Respuesta {
	var r Respuesta
	r.login = login
	r.nonce = getTranKey()
	r.seed = _seed
	r.tranKey = _nonce
	return r
}

func main() {
	generate()
	_nonce = getNonce()
	var a = asObject()
	fmt.Println(a)
}
