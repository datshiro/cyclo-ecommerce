#!/usr/bin/env bash -l

. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"
. "$(git rev-parse --show-toplevel || echo ".")/api/.env"

cmd=$1


get_one(){
    http -v GET $API_URL/products/${1:-"1"}
}

get_many(){
    http -v GET $API_URL/products  min_price:=${2:-50} max_price:=${3:-70} sort_by=${4:-"brands"}
}

"$@"

