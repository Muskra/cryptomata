#!/bin/bash

value="t3ZtP4sW0rDr34lYZeKvr3"

key1="o5eQBuW8npe26oS4dSpY2QaJ9zd8sYuliKck5OeTLeLTFPbnuIi3ZHAln5zAMTI7xl9bKlEXBfd9edhAddixJlSSMNHE8QQDpMtbglLJ7BZE92bewUq2mbgJ21BAS3wk"

key2="InHLG7TaqO4cdocfIF4K9pMaox99m7IUbB99HMvCYDEiQQRtSssC4pqwdqgPbSNGShWiEhgAoWPHOD8zAfFIwS2NvNribbqOUkIXljTFycThMRniZDlc4vybXUuo0JjS"

first=$(./cryptomata $value $key1 $key2)
second=$(./cryptomata $first $key2 $key1)

echo "input : $value"
echo "encryption : $first"
echo "decryption : $second"
