package main

import (
   "log"
   "./utils"
)

func main() {
   key := "111023043350789514532147"
   message := "Iam A message"
   log.Println("Original message: ", message)
   encryptedString := utils.EncryptString(key, message)
   log.Println("Encrypted message: ", encryptedString)
   decryptedString := utils.DecryptString(key, encryptedString)
   log.Println("Decrypted message: ", decryptedString)
}

