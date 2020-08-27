type Request struct {
   Method string // GET, POST, PUT, etc
   URL *url.URL // Parsed URL
   Proto string // "HTTP/1.0"
   ProtoMajor int // 1
   ProtoMinor int // 0
   // A header maps request lines to their values
   Header Header // map[string][]string
   
   // The message body.
   Body io.ReadCloser
   
   // ContentLength records the length of the associated content.
   // The value -1 indicates that length is unknown
   // values >= 0 indicate that the given number of bytes may be 
   // read from Body.
   ContentLength int64
   
   // TransferEncoding lists the transfer encodings from outermost to 
   // innermost. An empty list denotes the "identity" encoding.
   TransferEncoding []string
   
   // The host on which the URL is sought
   // Per RFC 2616, this is either the value of the Host: header
   // or the host name given in the URL itself.
   Host string
}
