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
	IpfsNode() *core.IpfsNode                       // XXX temporary
	Cat(Path) (Data, error)                         // http GET
	Ls(Path) ([]Link, error)                        // http GET, PUT
	Add(Data) (Path, error)                         // http POST
	ObjectSetData(Path, Data) (Path, error)         // http PUT update
	ObjectAddLink(Path, string, Path) (Path, error) // http PUT create
	ObjectRmLink(Path, string, Path) (Path, error)  // http DELETE
	// PUT and DELETE only for subdirs: /ipfs/<hash>/foo ???
}

type Path string

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
