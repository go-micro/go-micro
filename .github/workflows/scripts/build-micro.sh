#!/bin/bash
# set -x

failed=0
go mod edit -replace github.com/micro/go-micro/v2=github.com/micro/go-micro/v2@$1 
# basic test, build the binary
go install
failed=$?
if [ $failed -gt 0 ]; then
    exit $failed
fi
# unit tests
IN_TRAVIS_CI=yes go test -v ./...

./scripts/test-docker.sh
pushd test
# Generate keys for JWT tests
ssh-keygen -f /tmp/sshkey -m pkcs8 -q -N ""
ssh-keygen -f /tmp/sshkey -e  -m pkcs8 > /tmp/sshkey.pub
IN_TRAVIS_CI=yes go clean -testcache && go test --tags=integration -v ./...
ret=$? 
popd
echo $ret