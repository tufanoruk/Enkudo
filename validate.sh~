#!/bin/bash
#
# Validate OpeAPI doccument usig variour validators
#


Usage="\
Usage: validate.sh <OpenAPI_formatted_file> \n
\tJSON/YAML should be OpenAPI 3.0.0 foramat \n
"

if [ ! $1 ]
then 
    echo -e $Usage
    exit
fi

INF=$1

#echo "===== Validating with openapi-generator ====="
#openapi-generator validate -i $INF

#echo "===== Validating with openapi-spec-validator ====="
#openapi-spec-validator $INF

VALIDATORS=("openapi-generator validate -i" 
            "openapi-spec-validator ")

for ((i = 0; i < ${#VALIDATORS[@]}; i++))
do
    echo "===== Validating with ${VALIDATORS[$i]} ====="
    ${VALIDATORS[$i]} ${INF}
    if [ $? -ne 0 ] 
    then
        echo "Invalid ${INF} file"
        exit -1
    fi
done
