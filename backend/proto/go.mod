module github.com/t-ash0410/tdd-sample/backend/proto

go 1.17

replace github.com/t-ash0410/tdd-sample/backend/proto => ./

replace github.com/t-ash0410/tdd-sample/backend/internal => ../internal

replace github.com/t-ash0410/tdd-sample/backend/pkg => ../pkg

replace github.com/t-ash0410/tdd-sample/backend/test => ../test

require (
	github.com/golang/protobuf v1.5.2
	github.com/pkg/errors v0.9.1
	github.com/t-ash0410/tdd-sample/backend/internal v0.0.0-00010101000000-000000000000
	github.com/t-ash0410/tdd-sample/backend/test v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

require (
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f // indirect
	golang.org/x/sys v0.0.0-20210906170528-6f6e22806c34 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210917145530-b395a37504d4 // indirect
)
