# fsck

`fsck` is a simple tool for verifying the integrity of a [`tlog-tiles`][] log.

It is so-named as a nod towards the 'nix tools which perform a similar job for filesystems.
Note, however, that this tool is generally applicable for all tlog-tile instances accessible
via a HTTP, not just those which _happen_ to be backed by a POSIX filesystem.

## Usage

The tool is provided the URL of the log to check, and will attempt to re-derive 
the claimed root hash from the log's `checkpoint`, as well as the contents of all
tiles implied by the tree size it contains.

It can be run with the following command:

```bash
$ go run github.com/transparency-dev/trillian-tessera/cmd/experimental/fsck --storage_url=http://localhost:2024/ 
I0515 11:53:10.652868  241971 fsck.go:54] Fsck: checkpoint:
TestTessera
193446
3dFT/ML7vbDp84UT+SVR+Y9csuztBzvu5yuZXoV4E+k=

— TestTessera V9zVSATPw3zQN8jQ5hVhDTqQv0IMov4Ax+2xYE0yUNyySVHfX5xxuka0/HiwjaWqI96ux4/5kZsEqLjUFCMnzX/z2AA=
I0515 11:53:10.652964  241971 fsck.go:60] Fsck: checking log of size 193446
I0515 11:53:10.653001  241971 stream.go:90] StreamAdaptor: streaming from 0 to 193446
I0515 11:53:11.297305  241971 fsck.go:118] Successfully fsck'd log with size 193446 and root ddd153fcc2fbbdb0e9f38513f92551f98f5cb2eced073beee72b995e857813e9
```

Optional flags may be used to control the amount of parallelism used during the process, run the tool with `--help`
for more details.

