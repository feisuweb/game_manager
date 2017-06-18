#! /bin/bash
echo 'build for linux'
bee pack -be GOOS=linux -exp=.:upload:logs:bee.json:README.md:shell  
