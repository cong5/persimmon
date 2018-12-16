package myxmlrpc

import (
	"fmt"
)

var (
	FaultInvalidParams        = Fault{Code: -32602, String: "Invalid Method Parameters"}
	FaultWrongArgumentsNumber = Fault{Code: -32602, String: "Wrong Arguments Number"}
	FaultInternalError        = Fault{Code: -32603, String: "Internal Server Error"}
	FaultApplicationError     = Fault{Code: -32500, String: "Application Error"}
	FaultSystemError          = Fault{Code: -32400, String: "System Error"}
	FaultDecode               = Fault{Code: -32700, String: "Parsing error: not well formed"}
)

type Fault struct {
	Code   int    `xml:"faultCode"`
	String string `xml:"faultString"`
}

func (f Fault) Error() string {
	return fmt.Sprintf("%d: %s", f.Code, f.String)
}
