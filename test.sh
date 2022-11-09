go test ./domains/... -coverprofile=cover.out && go tool cover -html=cover.out
go tool cover -func cover.out