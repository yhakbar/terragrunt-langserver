#!/usr/bin/env bash

CODE_TESTS_PATH="$PWD/client/out/test"
export CODE_TESTS_PATH
CODE_TESTS_WORKSPACE="$PWD/client/testFixture"
export CODE_TESTS_WORKSPACE

node "$PWD/client/out/test/runTest"
