#!/bin/bash

 for file in `ls |grep abc`;do mv "$file" `echo "$file" |sed 's/abc/cba/g'`;done

