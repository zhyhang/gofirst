package dbconst

const (
	MsgDsn = `Please specifiy connection parameter in GO_OCI8_CONNECT_STRING environment variable,
or as the first argument! (The format is user/pass@host:port/sid)`
	DefaultDsn = "sys/sys@127.0.0.1/orcltest?as=sysdba"
)
