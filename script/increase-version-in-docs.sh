#!/bin/bash

grep -rl "0\.3\.0" example | xargs sed -i "" -e "s/0\.3\.0/0.3.1/g"
#grep -rl "0\.3\.0" example
grep -rl "0\.3\.0" docs | xargs sed -i "" -e "s/0\.3\.0/0.3.1/g"
#grep -rl "0\.3\.0" docs
