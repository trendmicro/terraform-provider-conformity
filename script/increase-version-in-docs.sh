#!/bin/bash

grep -rl "0\.3\.4" example | xargs sed -i "" -e "s/0\.3\.4/0.3.5/g"
grep -rl "0\.3\.4" docs | xargs sed -i "" -e "s/0\.3\.4/0.3.5/g"
