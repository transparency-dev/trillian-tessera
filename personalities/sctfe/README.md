# SCTFE

This personality implements https://c2sp.org/static-ct-api using
Trillian Tessera to store data. It is based on [Trillian's CTFE](https://github.com/google/certificate-transparency-go/tree/master/trillian/ctfe).

It is under active development, tracked under [Issue#88](https://github.com/transparency-dev/trillian-tessera/issues/88).

## Deployment
Each Tessera storage backend needs its own SCTFE binary.

At the moment, these storage backends are supported:

 - [GCP](./ct_server_gcp)


TODO(phbnf): add deployment instructions


## Working on the Code
The following files are auto-generated:
 - [`config.pb.go`](./configpb/config.pb.go): SCTFE's config
 - [`mock_ct_storage.go`](./mockstorage/mock_ct_storage.go): a mock CT storage implementation for tests

To re-generate these files, first install the right tools:
 - [protobuf compiler and go gen](https://protobuf.dev/getting-started/gotutorial/#compiling-protocol-buffers). The protos in this repo have been built with protoc v27.3.
 - [mockgen](https://github.com/golang/mock?tab=readme-ov-file#installation)

Then, generate the files:
```bash
cd $(go list -f '{{ .Dir }}' github.com/transparency-dev/trillian-tessera/personalities/sctfe); \
go generate -x ./...  # hunts for //go:generate comments and runs them
```

TODO(phboneff): provide docker template to build everything
