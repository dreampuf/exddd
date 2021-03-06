#!/bin/bash

PACKAGE_NAME=web/web.tar.gz
LOCAL_PATH=local_package
TARGET_PATH=/opt/exddd

mkdir $LOCAL_PATH
tar -xf $PACKAGE_NAME -C $LOCAL_PATH
cp system_config/{exddd.nginx.conf,exddd.supervisor.conf} $LOCAL_PATH
rsync -av $LOCAL_PATH/ root@$TARGET_HOST:$TARGET_PATH/
ssh root@$TARGET_HOST "ln -snf $TARGET_PATH/exddd.nginx.conf /etc/nginx/sites-available/exddd.nginx.conf"
ssh root@$TARGET_HOST "ln -snf /etc/nginx/sites-available/exddd.nginx.conf /etc/nginx/sites-enabled/exddd.nginx.conf"
ssh root@$TARGET_HOST "pkill -f web || true"
ssh root@$TARGET_HOST "nohup /opt/exddd/run.sh > /dev/null &"
