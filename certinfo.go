package certinfo

import (
   "bytes"
   "crypto/dsa"
   "crypto/ecdsa"
   "crypto/rsa"
   "crypto/x509"
   "crypto/x509/pkix"
   "encoding/asn1"
   "errors"
   "fmt"
   "math/big"
   "net"
   "time"
)

// Extra ASN1 OID's that we may need to handle
var (
   oidEmailAddress                 = []int{1, 2, 840, 113549, 1, 9, 1}
   oldExtensionAuthorityInfoAccess = []int{1, 3, 6, 1, 5, 5, 7, 1, 1}
   oidNSComment                    = []int{2, 16, 840, 1, 113730, 1, 13}
)
