#!/bin/bash
if [ ! -d "bin" ]; then
  mkdir bin
fi

go build cmd/server.go
mv server bin/
if [ $? -eq 0 ]; then
  echo -e "------ \033[32m[succ]\033[0m go build server: OK ------"
else
  echo -e "------ \033[31m[fail]\033[0m go build server: NG ------"
  exit -1;
fi

