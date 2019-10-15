#!/bin/sh
#curl -s 'https://api.github.com/users/diliska' | -r '.id'
curl -s 'https://api.github.com/users/diliska' | jq -r '.id'