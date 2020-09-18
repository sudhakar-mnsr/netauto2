#!/bin/bash
exec 100<>/dev/tcp/hello.natwelch.com/80
echo -e "GET / HTTP/1.1\r\nHost: hello.natwelch.com\r\nConnection: close\r\n\r\n" >&100
cat <&100
