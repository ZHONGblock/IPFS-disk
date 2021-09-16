package ipfs

import (
	"io"

	shell "github.com/ipfs/go-ipfs-api"
)

func AddFile(r io.Reader, options ...shell.AddOpts) (string, error) {
	sh := shell.NewShell("localhost:ipfsport")
	cid, err := sh.Add(r)
	if err != nil {
		return "", err
	}

	return cid, err
}
