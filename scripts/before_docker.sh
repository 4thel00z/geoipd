#! /bin/sh

this_dir=$(dirname "$0")
export PREFIX=/usr/local
mkdir -p $PREFIX/go/src/geoipd $PREFIX/go/src/_/builds
cp -r $this_dir/../* $PREFIX/go/src/geoipd
ln -s $PREFIX/go/src/geoipd $PREFIX/go/src/_/builds/geoipd
