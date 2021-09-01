#!/bin/bash

grep -rl "0\.3\.2" example | xargs sed -i "" -e "s/0\.3\.2/0.3.3/g"
grep -rl "0\.3\.2" docs | xargs sed -i "" -e "s/0\.3\.2/0.3.3/g"
