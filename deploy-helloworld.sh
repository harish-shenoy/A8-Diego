#!/bin/bash

A8_VERSION=0.4.2
A8_SRC=src/github.com/amalgam8
PROJ_HOME=$PWD
DEPLOY=$PROJ_HOME/deploy


# Backup parent process GO env vars so they can be restored when build is complete
GOHOSTOS_BACKUP=$GOHOSTOS
GOHOSTARCH_BACKUP=$GOHOSTARCH
GOPATH_BACKUP=$GOPATH

# Set GO env vars for CF stack
export GOHOSTOS=linux
export GOHOSTARCH=amd64
export GOPATH=$PROJ_HOME/sandbox

echo Downloading amalgam8-v$A8_VERSION
#curl -L -o amalgam8.zip https://github.com/amalgam8/amalgam8/archive/v$A8_VERSION.zip

echo Creating directory structure $GOPATH/$A8_SRC
mkdir -p $GOPATH/$A8_SRC
mkdir -p $DEPLOY

echo Unzip archive to $GOPATH/$A8_SRC
unzip amalgam8.zip -d $GOPATH/$A8_SRC

echo Rename amalgam8 source directory
mv $GOPATH/$A8_SRC/amalgam8-$A8_VERSION $GOPATH/$A8_SRC/amalgam8

cd $GOPATH/$A8_SRC/amalgam8

make build.sidecar

ls -al bin/

mv bin/* $DEPLOY
cp controller/schema.json $DEPLOY

cd $PROJ_HOME

# Restore parent space GOPATH
export GOPATH=$GOPATH_BACKUP
export GOHOSTOS=$GOHOSTOS_BACKUP
export GOHOSTARCH=$GOHOSTARCH_BACKUP
