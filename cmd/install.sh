#! /usr/bin/bash
THISDIR=$(dirname $(readlink -f $0))

for I in $THISDIR/*
do
  if [ -d $I ]; then
    go install $I
fi
done

