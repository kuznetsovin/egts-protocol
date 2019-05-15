#/bin/bash

LIB_PATH=./bin/plugins

if [ "$1" != "" ]; then
    mkdir -p $LIB_PATH
    go build -buildmode=plugin -o $LIB_PATH/$1.so ../pkg/store-plugins/$1/$1.go
else
    echo "Не задано имя плагина"
fi