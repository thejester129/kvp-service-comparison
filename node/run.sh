#!/bin/bash

cd "${0%/*}" # ensure cwd is script dir

# node 24.0
node --watch --experimental-transform-types src/index.ts 