// const.go
package main

// AUTH_METHOD
const (
	NOAUTH               = iota
	GSSAPI               = iota
	USERNAME_PASSWORD    = iota
	NO_ACCEPTABLE_METHOD = 0xFF
)

// COMMAND
const (
	CONNECT       = 1 + iota
	BIND          = 1 + iota
	UDP_ASSOCIATE = 1 + iota
)

// ADDRTYPE
const (
	IPV4   = 1 + iota
	DOMAIN = 2 + iota
	IPV6   = 2 + iota
)

// REQUEST_STATUS
const (
	SUCCESS                = iota
	SERVER_FAILURE         = iota
	CONNECTION_NOT_ALLOWED = iota
	NETWORK_UNREACHABLE    = iota
	HOST_UNREACHABLE       = iota
	CONNECTION_REFUSED     = iota
	TTL_EXPIRED            = iota
	COMMAND_NOT_SUPPORTED  = iota
	ADDRTYPE_NOT_SUPPORTED = iota
)
