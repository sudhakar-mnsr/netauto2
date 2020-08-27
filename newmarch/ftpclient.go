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

func main() {
   if len(os.Args) != 2 {
      fmt.Println("Usage: ", os.Args[0], "host")
      os.Exit(1)
   }
   
   host := os.Args[1]
   
   conn, err := net.Dial("tcp", host+":1202")
   checkError(err)
   
   reader := bufio.NewReader(os.Stdin)
   for {
      line, err := reader.ReadString('\n')
      // lose trailing whitespace
      line = strings.TrimRight(line, " \t\r\n")
      if err != nil {
         break
      }
      
      // split into command + arg
      strs := strings.SplitN(line, " ", 2)
      // decode user request
      switch strs[0] {
      case uiDir:
         dirRequest(conn)
      case uiCd:
         if len(strs) != 2 {
            fmt.Println("cd <dir>)
            continue
         }
         fmt.Println("CD \"",strs[1], "\"")
         cdRequest(conn, strs[1])
      case uiPwd:
         pwdRequest(conn)
      case uiQuit:
         conn.Close()
         os.Exit(0)
      default:
         fmt.Println("Unknown command")
      }
   }
}
