#!/bin/bash

#A8_VERSION=0.4.2
A8_VERSION=$1
if [ "$A8_VERSION" = "master" ]
then
  A8_ZIP=$A8_VERSION.zip
else
  A8_ZIP=v$A8_VERSION.zip
fi
A8_SRC=src/github.com/amalgam8
PROJ_HOME=$PWD
DEPLOY=$PROJ_HOME/deploy

# Set GO env vars for CF stack
GOHOSTOS=linux
GOHOSTARCH=amd64
GOPATH=$PROJ_HOME/sandbox

echo Downloading amalgam8-$A8_VERSION
curl -L -o amalgam8.zip https://github.com/amalgam8/amalgam8/archive/$A8_ZIP

echo Creating directory structure $GOPATH/$A8_SRC
mkdir -p $GOPATH/$A8_SRC
mkdir -p $DEPLOY

echo Unzip archive to $GOPATH/$A8_SRC
unzip amalgam8.zip -d $GOPATH/$A8_SRC

echo Rename amalgam8 source directory
mv $GOPATH/$A8_SRC/amalgam8-$A8_VERSION $GOPATH/$A8_SRC/amalgam8

cd $GOPATH/$A8_SRC/amalgam8

GOPATH=$GOPATH GOOS=$GOHOSTOS GOARCH=$GOHOSTARCH make build.controller build.registry build.sidecar

ls -al bin/

mv bin/* $DEPLOY
cp controller/schema.json $DEPLOY

cd $PROJ_HOME

cf push a8controller-02 -f manifest-binary.yml
cf push a8registry-02 -f manifest-binary.yml
