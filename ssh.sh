#!/bin/bash 
/usr/bin/ssh -A -l $1 $2 | tee $3
ssh-add -D
