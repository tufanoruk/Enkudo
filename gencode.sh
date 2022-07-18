#!/bin/bash
#

generator="openapi-generator"

Usage="\
Usage: gencode.sh <YAML_specification> <Language> \n
\tYAML should be OpenAPI 3.0.0 foramat \n
\tLanguage should be one of 'openapi-generator list'\n 
"

if [ ! $2 ]
then 
    echo -e $Usage
    exit
fi

SPEC=$1
LANG=$2

echo "Generating $LANG code from $SPEC"

$generator validate -i $SPEC
if [ $? -ne 0 ] 
then
    echo "Invalid $SPEC file"
    exit
fi

SRC=$LANG"_src"

$generator generate \
    --generator-name $LANG \
    --api-name-suffix Enkudo_DSP \
    --output $SRC \
    --input-spec $SPEC
