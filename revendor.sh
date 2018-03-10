#!/bin/bash
set -euo pipefail
mkdir .tmp

git clone https://github.com/google/go-cmp/ .tmp/go-cmp
cp -r .tmp/go-cmp/cmp/ .
cp .tmp/go-cmp/LICENSE cmp
find cmp/ -type f -name \*.go -print0 | xargs -0 sed -i "s#github.com/google/go-cmp/cmp#github.com/warpfork/go-wish/cmp#"


git clone https://github.com/pmezard/go-difflib .tmp/go-difflib
cp -r .tmp/go-difflib/difflib/ .
cp .tmp/go-difflib/LICENSE difflib
