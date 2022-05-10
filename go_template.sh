#!/usr/bin/bash -eu

script_dir="$(cd $(dirname $(readlink -f $0)); pwd)"

if [ $# != 1 ]; then
    echo "引数が不正です"
    exit 1
fi

if [ -d $1 ]; then
    if [ ! -z "$(ls -A $1)" ]; then
        echo "フォルダはすでに存在します"
        exit 1
    fi
else
    echo "creating folder: $1"

    mkdir $1
    if [ $? != 0 ]; then
        echo "フォルダの作成に失敗しました"
        exit 1
    fi
fi

target_dir="$(readlink -f $1)"

cd $target_dir

go mod init $1

cd "$script_dir/dir/"
ls -A | xargs -I {} cp -r {} $target_dir

cd $target_dir
code -n . -g "main.go":51:2
