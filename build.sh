#!/bin/bash
# Thanks to https://github.com/bettercap/bettercap/build.sh by @evilsocket

BUILD_FOLDER=build
VERSION=v$(cat cmd/root.go | grep Version | cut -d '"' -f 2)

create_exe_archive() {
    OUTPUT=$1

    echo "@ Creating archive $OUTPUT ..."
    zip -j "$OUTPUT" hosty.exe ../README.md ../LICENSE > /dev/null
    rm -rf hosty.exe
}

create_archive() {
    OUTPUT=$1

    echo "@ Creating archive $OUTPUT ..."
    zip -j "$OUTPUT" hosty ../README.md ../LICENSE > /dev/null
    rm -rf hosty
}

build_linux_amd64() {
    echo "@ Building linux/amd64 ..."
    GOOS="linux" GOARCH="amd64" go build -o hosty ..
}

build_windows_amd64() {
    echo "@ Building windows/amd64"
    GOOS="windows" GOARCH="amd64" go build -o hosty.exe ..
}

build_darwin_amd64() {
    echo "@Building darwin/amd64"
    GOOS="darwin" GOARCH="amd64" go build -o hosty ..
}

# cleanup build environment
rm -rf $BUILD_FOLDER
mkdir $BUILD_FOLDER
cd $BUILD_FOLDER

# check arguments
if [ -z "$1" ]
    then
        WHAT=all
    else
        WHAT="$1"
fi

printf "@ Building for $WHAT ...\n\n"

if [[ "$WHAT" == "all" || "$WHAT" == "linux_amd64" ]]; then
    build_linux_amd64 && create_archive hosty_linux_amd64_$VERSION.zip
    sha256sum ./hosty_linux_amd64_$VERSION.zip > hosty_linux_amd64_$VERSION.sha256
fi

if [[ "$WHAT" == "all" || "$WHAT" == "darwin_amd64" ]]; then
    build_darwin_amd64 && create_archive hosty_darwin_amd64_$VERSION.zip
    sha256sum ./hosty_darwin_amd64_$VERSION.zip > hosty_darwin_amd64_$VERSION.sha256
fi

if [[ "$WHAT" == "all" || "$WHAT" == "windows" || "$WHAT" == "win" ]]; then
    build_windows_amd64 && create_exe_archive hosty_windows_amd64_$VERSION.zip
    sha256sum ./hosty_windows_amd64_$VERSION.zip > hosty_windows_amd64_$VERSION.sha256
fi

du -sh *

cd ..
