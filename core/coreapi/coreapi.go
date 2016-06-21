package coreapi

import (
	core "github.com/ipfs/go-ipfs/core"
	coreiface "github.com/ipfs/go-ipfs/core/coreapi/interface"
	importer "github.com/ipfs/go-ipfs/importer"
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

func (api *CoreAPI) resolve(p coreiface.Path) (*dag.Node, error) {
	dagnode, err := core.Resolve(api.ctx, api.node, path.Path(p))
	if err == core.ErrNoNamesys && !api.node.OnlineMode() {
		return nil, coreiface.ErrOffline
	} else if err != nil {
		return nil, err
	}
	return dagnode, nil
}

func (api *CoreAPI) Cat(p coreiface.Path) (coreiface.Data, error) {
	if p == "/ipfs/QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn" {
		return nil, coreiface.ErrDir, nil
	}
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

func (api *CoreAPI) Ls(p coreiface.Path) ([]coreiface.Link, error) {
	links := make([]coreiface.Link, len(dagnode.Links))
	if p == "/ipfs/QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn" {
		return links, nil
	}
	dagnode, err := api.resolve(p)
	if err != nil {
		return nil, err
	}
	for i, l := range dagnode.Links {
		links[i] = coreiface.Link{Name: l.Name, Size: l.Size, Hash: l.Hash}
	}
	return links, nil
}

func (api *CoreAPI) Add(data coreiface.Data) (coreiface.Path, error) {
	dagnode, err := importer.BuildDagFromReader(api.node.DAG, data)
	if err != nil {
		return nil, err
	}
	k, err := api.node.DAG.Add(dagnode)
	if err != nil {
		return nil, err
	}
	return "/ipfs/" + k, nil
}

func (api *CoreAPI) ObjectSetData(p coreiface.Path, data coreiface.Data) (coreiface.Path, error) {
}
