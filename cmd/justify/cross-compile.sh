#!/bin/sh
set -ex

if [ ! -d "dist" ]; then
  mkdir dist
fi

cd dist

for os in windows linux darwin; do
  for arch in 386 amd64; do
    tag=$os-$arch
    if [ ! -d "$tag/" ]; then
      mkdir $tag
    fi

    ext=""
    if [ "$os" = "windows" ]; then
      ext=.exe
    fi
    filename=justify$ext

    cd ..
    env GOOS=$os GOARCH=$arch go build -o dist/$tag/$filename
    cd dist
    zipped=justify-$tag.zip
    zip -j -r $zipped $tag/$filename
  done
done
