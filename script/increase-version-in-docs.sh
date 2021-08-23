#!/bin/bash

grep -rl "0\.3\.1" example | xargs sed -i "" -e "s/0\.3\.1/0.3.2/g"
grep -rl "0\.3\.1" docs | xargs sed -i "" -e "s/0\.3\.1/0.3.2/g"
