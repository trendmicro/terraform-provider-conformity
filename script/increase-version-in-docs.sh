#!/bin/bash

grep -rl "0\.2\.0" example | xargs sed -i "" -e "s/0\.2\.0/0.3.0/g"
grep -rl "0\.2\.0" docs | xargs sed -i "" -e "s/0\.2\.0/0.3.0/g"
