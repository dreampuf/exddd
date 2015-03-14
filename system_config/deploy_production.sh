#!/bin/bash

PACKAGE_NAME=web/web.tar.gz
LOCAL_PATH=local_package

mkdir $LOCAL_PATH
tar -xf $PACKAGE_NAME -C $LOCAL_PATH
rsync -av $LOCAL_PATH/ root@$TARGET_HOST:/opt/exddd/
