#!/bin/bash
set -e
 
count=$(find . | wc -l)
result=$((count*5))
printf "\t\vTotal files * 5: $result\v\n"
