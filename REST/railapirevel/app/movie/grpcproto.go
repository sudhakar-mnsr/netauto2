syntax = "proto3"
package datafiles;

message TransactionRequest {
   string from = 1;
   string to = 2;
   float amount = 3;
}
