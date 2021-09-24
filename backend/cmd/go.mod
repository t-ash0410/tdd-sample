module github.com/t-ash0410/tdd-sample/backend/cmd

go 1.17

replace github.com/t-ash0410/tdd-sample/backend/proto => ../proto

replace github.com/t-ash0410/tdd-sample/backend/internal => ../internal

replace github.com/t-ash0410/tdd-sample/backend/pkg => ../pkg

replace github.com/t-ash0410/tdd-sample/backend/test => ../test

require (
	github.com/t-ash0410/tdd-sample/backend/internal v0.0.0-00010101000000-000000000000
	github.com/t-ash0410/tdd-sample/backend/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.40.0
)

require (
	cloud.google.com/go v0.94.1 // indirect
	cloud.google.com/go/spanner v1.25.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/googleapis/gax-go/v2 v2.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/t-ash0410/tdd-sample/backend/pkg v0.0.0-00010101000000-000000000000 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f // indirect
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f // indirect
	golang.org/x/sys v0.0.0-20210906170528-6f6e22806c34 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/api v0.56.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210917145530-b395a37504d4 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
