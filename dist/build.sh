#!/bin/bash
set -e

# This script should not be verbose.
# Simply telling what it is doing is enough.
GITVERSION="+git"$(date +%Y%m%d%H%M)"."$(git log -n 1 | tr " " "\n" | head -2 | tail -1 | head -c 7)
function ok {
    echo "OK"
}

root=$(dirname $0)
cd "$root"
root=$(pwd)
vcode="1.0.2-"$(cat ../VERSION_CODE | head -1)
echo "Building dostestportal - version: $vcode";
cd ..
rm -rf build/
goprodbuilds=$(pwd)"/build/"
~/go/bin/packr2 clean
~/go/bin/packr2
goprod \
    -combo="linux/arm;linux/386;linux/arm64;linux/amd64;windows/amd64;windows/386" \
    -binname="dostestportal" \
    -tags="guibrowser" \
    -version="$vcode"
~/go/bin/packr2 clean