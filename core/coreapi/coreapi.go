package coreapi

import (
	core "github.com/ipfs/go-ipfs/core"
	coreiface "github.com/ipfs/go-ipfs/core/coreapi/interface"
	dag "github.com/ipfs/go-ipfs/merkledag"
	path "github.com/ipfs/go-ipfs/path"
	uio "github.com/ipfs/go-ipfs/unixfs/io"
	context "gx/ipfs/QmZy2y8t9zQH2a1b8q2ZSLKp17ATuJoCNxxyMFG5qFExpt/go-net/context"
)

type CoreAPI struct {
	ctx  context.Context
	node *core.IpfsNode
}

func NewCoreAPI(ctx context.Context, node *core.IpfsNode) (*CoreAPI, error) {
	api := &CoreAPI{ctx: ctx, node: node}
	return api, nil
}

func (api *CoreAPI) Context() context.Context {
	return api.ctx
}

func (api *CoreAPI) IpfsNode() *core.IpfsNode {
	return api.node
}

func (api *CoreAPI) resolve(p string) (*dag.Node, error) {
	dagnode, err := core.Resolve(api.ctx, api.node, path.Path(p))
	if err == core.ErrNoNamesys && !api.node.OnlineMode() {
		return nil, coreiface.ErrOffline
	} else if err != nil {
		return nil, err
	}
	return dagnode, nil
}

func (api *CoreAPI) Cat(p string) (coreiface.Data, error) {
	dagnode, err := api.resolve(p)
	if err != nil {
		return nil, err
	}
	r, err := uio.NewDagReader(api.ctx, dagnode, api.node.DAG)
	if err == uio.ErrIsDir {
		return nil, coreiface.ErrDir
	} else if err != nil {
		return nil, err
	}
	return r, nil
}

func (api *CoreAPI) Ls(p string) ([]coreiface.Link, error) {
	dagnode, err := api.resolve(p)
	if err != nil {
		return nil, err
	}
	links := make([]coreiface.Link, len(dagnode.Links))
	for i, l := range dagnode.Links {
		links[i] = coreiface.Link{Name: l.Name, Size: l.Size, Hash: l.Hash}
	}
	return links, nil
}
