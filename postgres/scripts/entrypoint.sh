#!/bin/bash -e

echo "setting up various databases"
pushd /home/postgres/migrations/
for d in */ ; do
    echo "setting up migrations in $d"
    pushd $d
    source setup.sh
    popd
done

