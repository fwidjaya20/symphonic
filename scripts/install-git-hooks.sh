#!/bin/bash
set -e

[ -d .githooks ] || mkdir .githooks
git config core.hooksPath .githooks
