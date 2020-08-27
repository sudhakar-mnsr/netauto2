/* FTPClient
*/
package main

import (
   "bufio"
   "bytes"
   "fmt"
   "net"
   "os"
   "strings"
)

// strings used by the user interface
const (
   uiDir = "dir"
   uiCd = "cd"
   uiPwd = "pwd"
   uiQuit = "quit"
)

// strings used across network
const (
   DIR = "DIR"
   CD = "CD"
   PWD = "PWD"
)
