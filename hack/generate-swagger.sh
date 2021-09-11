#! /bin/bash
# This script is used to generate swagger documentation.
# use `swagger serve SWAGGER_API_JSON` to launch swagger server

# check command tool
if ! command -v protoc > /dev/null
then
    echo -e "\033[31mprotoc could not be found\033[0m"
    exit
fi
if ! command -v swagger > /dev/null
then
    echo -e "\033[31mswagger could not be found\033[0m"
    exit
fi


REPO_DIR=$1
GRPC_GATEWAY_PROTOCOL=$2
KRATOS=$3

find $REPO_DIR -path $REPO_DIR/configs -prune -o -name *.proto -print | xargs -n 1  protoc -I$GRPC_GATEWAY_PROTOCOL\
    --proto_path=$SCHEDULER_PROTO \
    --proto_path=$REPO_DIR \
    --proto_path=$KRATOS/third_party \
    --openapiv2_out $REPO_DIR \
    --openapiv2_opt logtostderr=true
mkdir -p $REPO_DIR/docs/api
find $REPO_DIR/api $REPO_DIR/docs/metadata -iname \*swagger.json -exec mv  {} $REPO_DIR/docs/api \;
swagger mixin $REPO_DIR/docs/api/*.json -q --ignore-conflicts -o $REPO_DIR/docs/apidocs.swagger.json
rm -rf $REPO_DIR/docs/api
echo use \'swagger serve $REPO_DIR/docs/apidocs.swagger.json\' to launch swagger server