package iface

import (
	"errors"
	"io"

	core "github.com/ipfs/go-ipfs/core"

	mh "gx/ipfs/QmYf7ng2hG5XBtJA3tN34DQ2GUN5HNksEw1rLDkmr6vGku/go-multihash"
	context "gx/ipfs/QmZy2y8t9zQH2a1b8q2ZSLKp17ATuJoCNxxyMFG5qFExpt/go-net/context"
)

type CoreAPI interface {
	Context() context.Context
	IpfsNode() *core.IpfsNode // XXX temporary
	Cat(string) (Data, error)
	Ls(string) ([]Link, error)
}

type Object struct {
	Links []Link
	Data  Data
}

type Link struct {
	Name string // utf-8
	Size uint64
	Hash mh.Multihash
}

type Data interface {
	io.Reader
	io.Seeker
	io.Closer
}

var ErrDir = errors.New("object is a directory")
var ErrOffline = errors.New("can't resolve, ipfs node is offline")
