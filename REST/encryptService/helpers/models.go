package helpers

// EncryptRequest structures request coming from client
type EncryptRequest struct {
   Text string `json:"text"`
   Key string `json:"key"`
}

// EncryptResponse structures response going to the client
type EncryptResponse struct {
   Message string `json:"message"`
   Err string `json:"error"`
}
