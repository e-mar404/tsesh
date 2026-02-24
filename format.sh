#!/usr/bin/env bash

if [ `gofmt -l . | wc -l` -gt 0 ]; then
  exit 1
fi
