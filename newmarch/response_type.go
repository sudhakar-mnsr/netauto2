type Response struct {
   Status string // e.g "200 OK"
   StatusCode int // e.g 200
   Proto string // e.g. "HTTP/1.0"
   ProtoMajor int // e.g. 1
   ProtoMinor int // e.g. 0
   
   Header map[string][]string
   
   Body io.ReadCloser
   
   ContentLength int64
   
   TransferEncoding []string
   
   Close bool
   
   Trailer map[string][]string
   
   Request *Request // the original request
   
   TLS *tls.ConnectionState // info about the TLS connection or nil
}
