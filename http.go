package plutogo

/*
#cgo pkg-config: plutobook

#include <plutobook.h>
#include <stdlib.h>
*/
import "C"

// SetSSLCAInfo Sets the path to a file containing trusted CA certificates
func SetSSLCAInfo(path string) {
	cPath := C.CString(path)
	C.plutobook_set_ssl_cainfo(cPath)
}

// SetSSLCAPath Sets the path to a directory containing trusted CA certificates
func SetSSLCAPath(path string) {
	cPath := C.CString(path)
	C.plutobook_set_ssl_capath(cPath)
}

// SetSSLVerifyPeer Enables or disables SSL peer certificate verification
func SetSSLVerifyPeer(verify bool) {
	cVerify := C.bool(verify)
	C.plutobook_set_ssl_verify_peer(cVerify)
}

// SetSSLVerifyHost Enables or disables SSL host name verification
func SetSSLVerifyHost(verify bool) {
	cVerify := C.bool(verify)
	C.plutobook_set_ssl_verify_host(cVerify)
}

// SetHttpFollowRedirects Enables or disables automatic following of HTTP redirects
func SetHttpFollowRedirects(follow bool) {
	cFollow := C.bool(follow)
	C.plutobook_set_http_follow_redirects(cFollow)
}

// SetHttpMaxRedirects Sets the maximum number of redirects to follow
func SetHttpMaxRedirects(amount int) {
	cAmount := C.int(amount)
	C.plutobook_set_http_max_redirects(cAmount)
}

// SetHttpTimeout Sets the maximum time allowed for an HTTP request
func SetHttpTimeout(timeoutSeconds int) {
	cSeconds := C.int(timeoutSeconds)
	C.plutobook_set_http_timeout(cSeconds)
}
