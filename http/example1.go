package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)


func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}
