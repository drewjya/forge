package template

func Env() string {
	return `ENV=development
PORT=5000
DATABASE_URL=postgresql://postgres:postgres@localhost:5432/postgres
JWT_SECRET=dev-secret
SSE_HEARTBEAT_SECONDS=10
JWT_EXPIRES_IN=168h
LOG_LEVEL=info # debug, info, warn, error
`
}
