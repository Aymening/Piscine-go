#!/bin/bash

data=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json)

name=$(echo "$data" | jq -r '.[] | select(.id == 70) | .name')

echo "\"$name\""

