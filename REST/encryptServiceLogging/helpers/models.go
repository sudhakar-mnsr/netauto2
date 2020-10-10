package helpers

// EncryptRequest structures request coming from client
type EncryptRequest struct {
   Text string `json:"text"`
   Key string `json:"key"`
}
