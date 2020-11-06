#!/bin/bash

OS="local"
while getopts "lh:" arg
do
    case $arg in

        h)  #help带all参数仅用于测试脚本，并备忘
            echo "-h [all] help"
            case $OPTARG in
                all)
                    echo "-l build linux bin, default local"
                    ;;
            esac
            exit
            ;;
        l)
            OS="linux"
            ;;
        ?)
            echo "unkonw argument"
            echo "-h all help"
            echo "-l build linux bin, default local"
            exit 1
            ;;
    esac
done

# Go Path
# CURDIR=`pwd`
# OLDGOPATH="$GOPATH"
# export GOPATH="$OLDGOPATH:$CURDIR"

LogPrefix=">>>>"

# 交叉编译
case  $OS  in
    linux)
        # Linux
        echo "$LogPrefix `date +"%H:%M:%S"` build linux bin"
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -installsuffix cgo -ldflags '-w'
        ;;
    *)
        # 本机
        echo "$LogPrefix `date +"%H:%M:%S"` build local bin"
        go build -mod=vendor -ldflags '-w'
        ;;
esac

echo -e "$LogPrefix `date +"%H:%M:%S"` \033[42;37m finished \033[0m"