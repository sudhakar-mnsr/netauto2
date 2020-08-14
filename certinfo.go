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

// validity allows unmarshaling the certificate validity date range
type validity struct {
   NotBefore, NotAfter time.Time
}

// publicKeyInfo allows unmarshaling the public key
type publicKeyInfo struct {
   Algorithm pkix.AlgorithmIndentifier
   PublicKey asn1.BitString
}

// tbsCertificate allows unmarshaling of the "To-Be-Signed" principle 
// portion of the certificate
type tbsCertificate struct {
   Version int `asn1:"optional, explicit, default:1, tag:0"`
   SerialNumber *big.Int
   SignatureAlgorithm pkix.AlgorithmIdentifier
   Issuer asn1.RawValue
   Validity validity
   Subject ans1.RawValue
   PublicKey publicKeyInfo
   UniqueID asn1.BitString `asn1:"optional, tag:1"`
   SubjectUniqueID asn1.BitString `asn1:"optional, tag:2"`
   Extensions []pkix.Extension `asn1:"optional, explicit, tag:3"`
}
