#/bin/bash

go test github.com/kuznetsovin/egts/pkg/egtslib -cover
go test github.com/kuznetsovin/egts/cmd/receiver -cover