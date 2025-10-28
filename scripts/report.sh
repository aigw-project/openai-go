#!/usr/bin/env bash

set -e
echo -n "" > coverage.txt

for d in $(go list .); do
    echo "--------Run test package: $d"
    GO111MODULE=on go test -gcflags="all=-N -l" -run="^(TestUnmarshal|TestForAIGW).*" -v -coverprofile=profile.out -covermode=atomic $d
    echo "--------Finish test package: $d"
    if [[ -f profile.out &&  $(cat profile.out| grep -v "^$"| wc -l) -gt 0 ]]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

