#!/bin/bash

# awk
cat ./file.txt | awk 'NR==10'
awk 'NR==10' file.txt

# sed 
cat ./file.txt | sed -n '10p'
sed -n "10p" file.txt

# tail 配合 head
cat ./file.txt | tail -n +10 | head -1
tail -n +10  file.txt | head -1
