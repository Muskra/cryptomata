#!/bin/bash

value="t3ZtP4sW0rDr34lYZeKvr3"
#value="THIS_ysasuPêùé13DKJ2UO8"
rd1=$(tr -dc A-Za-z0-9 </dev/urandom |head -c 1536; echo)
rd2=$(tr -dc A-Za-z0-9 </dev/urandom |head -c 2560; echo)
rd3=$(tr -dc A-Za-z0-9 </dev/urandom |head -c 1536; echo)
rd4=$(tr -dc A-Za-z0-9 </dev/urandom |head -c 2560; echo)

first=$(./cryptomata $value $rd1 $rd2 $rd3 $rd4)
second=$(./cryptomata $first $rd1 $rd2 $rd3 $rd4)

echo "input : $value"
echo "encryption : $first"
echo "decryption : $second"
