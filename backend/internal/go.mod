module github.com/t-ash0410/tdd-sample/backend/internal

go 1.16

replace github.com/t-ash0410/tdd-sample/backend/internal => ./

replace github.com/t-ash0410/tdd-sample/backend/pkg => ../pkg

replace github.com/t-ash0410/tdd-sample/backend/test => ../test

require (
	cloud.google.com/go/spanner v1.25.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/t-ash0410/tdd-sample/backend/pkg v0.0.0-00010101000000-000000000000
	github.com/t-ash0410/tdd-sample/backend/test v0.0.0-00010101000000-000000000000
	google.golang.org/api v0.56.0
	google.golang.org/genproto v0.0.0-20210917145530-b395a37504d4
	google.golang.org/grpc v1.40.0
)
