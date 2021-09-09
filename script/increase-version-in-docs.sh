#!/bin/bash

grep -rl "0\.3\.5" example | xargs sed -i "" -e "s/0\.3\.5/0.3.6/g"
grep -rl "0\.3\.5" docs | xargs sed -i "" -e "s/0\.3\.5/0.3.6/g"
