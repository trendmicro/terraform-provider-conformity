#!/bin/bash

grep -rl "0\.3\.3" example | xargs sed -i "" -e "s/0\.3\.3/0.3.4/g"
grep -rl "0\.3\.3" docs | xargs sed -i "" -e "s/0\.3\.3/0.3.4/g"
