#!/bin/bash
CURDIR=$(pwd)
for file in *.go
do
    echo "mockgen -source ${CURDIR}/${file} -destination=${CURDIR}/pmgomock/${file} -package pmgomock -imports \".=github.com/bukalapak/pmgo\""
    mockgen -source ${CURDIR}/${file} -destination=${CURDIR}/pmgomock/${file} -package pmgomock -imports ".=github.com/bukalapak/pmgo"
done
