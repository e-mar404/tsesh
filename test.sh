#!/usr/bin/env bash

go test ./... -v

if [ $? -ne 0 ]; then
  exit 1
fi
