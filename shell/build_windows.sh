#! /bin/bash
echo 'build for windows'
bee pack -be GOOS=windows -exp=.:upload:logs:bee.json:README.md:shell
