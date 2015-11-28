#!/bin/sh

GOFILE=$1
OUTPUTDIR=$2
FILE=$3

if [ "$GOFILE" = "" -o "$OUTPUTDIR" = "" -o "$FILE" = "" ]; then
   echo "Requires go source, output directory, and output filename as input."
   exit 1
fi

mkdir -p $OUTPUTDIR/darwin64
mkdir -p $OUTPUTDIR/lx-amd64
mkdir -p $OUTPUTDIR/arm6
mkdir -p $OUTPUTDIR/win32
mkdir -p $OUTPUTDIR/arm64

# cross-compile
GOOS=darwin GOARCH=amd64 go build -o $OUTPUTDIR/darwin64/$FILE $GOFILE
GOOS=linux GOARCH=amd64 go build -o $OUTPUTDIR/lx-amd64/$FILE $GOFILE
GOOS=linux GOARCH=arm go build -o $OUTPUTDIR/arm6/$FILE $GOFILE
GOOS=windows GOARCH=386 go build -o $OUTPUTDIR/win32/$FILE.exe $GOFILE
# Go 1.5 required for arm64 (ARM8)
GOOS=linux GOARCH=arm64 go build -o $OUTPUTDIR/arm64/$FILE $GOFILE

