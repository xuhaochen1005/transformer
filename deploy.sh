#!/bin/bash
set -e
if [ ! $1 ]; then
    echo "Usage: deploy [dev|pro|uat]"
    exit 0
fi
echo $1
executeFun() {
    echo "Step 1: building ..."
    mkdir -p deploy_tmp/transformer_mz/website
    cp config/"$1".yaml deploy_tmp/transformer_mz.yaml
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deploy_tmp/transformer_mz_bin server/server.go
    cd website && eval `cat .env.$1` npm run build:$1 && cd ../
    cp -r website/dist/* deploy_tmp/transformer_mz/website/
    tar -zcvf transformer_mz_deploy.tar.gz deploy_tmp
    echo "Step 2: Uploading ..."
    scp transformer_mz_deploy.tar.gz root@47.112.167.85:/tmp/transformer_mz_deploy.tar.gz
    ssh root@47.112.167.85  `sudo sh -c '
    INSTALL_DIR="/usr/local/transformer_mz" && echo "$INSTALL_DIR" \
    && mkdir -p "$INSTALL_DIR"_deploy && mkdir -p "$INSTALL_DIR" \
    && rm -rf "$INSTALL_DIR"_back \
    && mv "$INSTALL_DIR" "$INSTALL_DIR"_back \
    && tar -xzvf /tmp/transformer_mz_deploy.tar.gz -C "$INSTALL_DIR"_deploy \
    && rm -rf /tmp/transformer_mz_deploy.tar.gz \
    && mv "$INSTALL_DIR"_deploy/deploy_tmp/transformer_mz /usr/local \
    && mv "$INSTALL_DIR"_deploy/deploy_tmp/transformer_mz_bin /usr/local/bin/transformer_mz \
    && chown -R root:root /usr/local/bin/transformer_mz \
    && chmod -R 754 /usr/local/bin/transformer_mz \
    && mv "$INSTALL_DIR"_deploy/deploy_tmp/transformer_mz.yaml /usr/local/etc/transformer_mz.yaml \
    && rm -rf "$INSTALL_DIR"_deploy'`
    rm -rf deploy_tmp && rm -f server_{$1} && rm -f transformer_mz_deploy_tar.gz
    echo "Step 3: Restarting ..."
    ssh root@47.112.167.85 "sudo systemctl restart transformer_mz"
}
case $1 in
    "dev")
        echo "Step 1: building ..."
        CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist_server.exe server/server.go
        cd website && npm run build && cd ../
        echo "Step 2: Restarting ..."
        ./dist_server.exe
        cd website && npm run start
        ;;
    "pro")
        ENV="pro"
        executeFun "pro"
        ;;
    "uat")
        ENV="uat"
        executeFun "uat"
        ;;
    "*")
        echo "Usage: deploy [dev|pro|uat]"
        exit
        ;;
esac




#
#echo "Step 3: Restarting ..."
#ssh root@server "systemctl restart example"