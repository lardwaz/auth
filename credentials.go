package auth

import (
	"fmt"
	"strings"
)

// Credentials represents connection to a service having a protocol, identifier, password, host and port
type Credentials struct {
	Protocol   string
	Identifier string
	Password   string
	Host       string
	Port       string
}

// String returns a human friendly representation of the credentials
func (c Credentials) String() string {
	return fmt.Sprintf("Protocol: %s Identifier: %s Password: %s Host: %s Port: %s", c.Protocol, c.Identifier, c.Password, c.Host, c.Port)
}

// GetProtocol returns Protocol portion of the credentials
func (c Credentials) GetProtocol() string {
	return c.Protocol
}

// GetIdentifier returns Identifier portion of the credentials
func (c Credentials) GetIdentifier() string {
	return c.Identifier
}

// GetPassword returns Password portion of the credentials
func (c Credentials) GetPassword() string {
	return c.Password
}

// GetHost returns Host portion of the credentials
func (c Credentials) GetHost() string {
	return c.Host
}

// GetPort returns Port portion of the credentials
func (c Credentials) GetPort() string {
	return c.Port
}

// NewCredentials returns a connection from string in the format protocol://username:password@host:port
func NewCredentials(conn string, defaults ...*Credentials) *Credentials {
	var (
		credentials   *Credentials
		creds, server string
	)

	if len(defaults) == 1 {
		credentials = defaults[0]
	} else {
		credentials = new(Credentials)
	}

	if i := strings.Index(conn, "://"); i != -1 {
		credentials.Protocol = conn[:i]
		conn = conn[i+3:]
	}

	if i := strings.Index(conn, "@"); i == -1 { //no username password
		server = conn
	} else {
		creds = conn[:i]
		server = conn[i+1:]
	}

	if creds == "" {
		//do nothing
	} else if i := strings.Index(creds, ":"); i != -1 { //both id and password
		credentials.Identifier = creds[:i]
		credentials.Password = creds[i+1:]
	} else { //only id
		credentials.Identifier = creds
	}

	if server == "" {
		//do nothing
	} else if i := strings.Index(server, ":"); i != -1 { //both host and port
		credentials.Host = server[:i]
		credentials.Port = server[i+1:]
	} else { //only host
		credentials.Host = conn
	}

	return credentials
}
