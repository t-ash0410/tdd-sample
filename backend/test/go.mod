module github.com/t-ash0410/tdd-sample/backend/test

go 1.17

replace github.com/t-ash0410/tdd-sample/backend/internal => ../internal

replace github.com/t-ash0410/tdd-sample/backend/pkg => ../pkg

replace github.com/t-ash0410/tdd-sample/backend/test => ./

require github.com/t-ash0410/tdd-sample/backend/internal v0.0.0-00010101000000-000000000000
