#!/bin/bash

echo "Building non-fips image"
docker build . --progress=plain -t non-fips -f Dockerfile.incorrect
echo "Building fips image"
docker build . --progress=plain -t fips -f Dockerfile.correct
echo "Running non-fips image"
docker run non-fips
echo "Running fips image"
docker run fips
