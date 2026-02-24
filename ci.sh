#!/usr/bin/env bash

./test.sh
if [ $? -ne 0 ]; then
  echo "ci will not pass test, fix it"
  exit 1
fi

./format.sh
if [ $? -ne 0 ]; then
  echo "ci will not pass format, fix it"
  exit 1
fi

echo "everything is ok, free to push"
