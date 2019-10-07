#!/bin/sh

# Build the project
npm run build

# Create the static directory
rm -rf ../static/
mkdir ../static/
cp -R build/* ../static/

# cleanup
rm -rf ./build