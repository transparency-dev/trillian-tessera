#!/bin/bash
#
# build_log.sh is a script for building a small test log.

set -e # exit on any error codes from sub-commands

DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

export LOG=${DIR}/log/
export LOG_PRIVATE_KEY="PRIVATE+KEY+example.com/log/testdata+33d7b496+AeymY/SZAX0jZcJ8enZ5FY1Dz+wTML2yWSkK+9DSF3eg"
export LOG_PUBLIC_KEY="example.com/log/testdata+33d7b496+AeHTu4Q3hEIMHNqc6fASMsq3rKNx280NI+oO5xCFkkSx"

cd ${DIR}
rm -fr log

go run ../cmd/posix-oneshot --storage_dir=${LOG} --initialise
cp ${LOG}/checkpoint ${LOG}/checkpoint.0

export LEAF=`mktemp`
for i in one two three four five six seven eit nain ten ileven twelf threeten fourten fivten; do
  echo -n "$i" > ${LEAF}
  go run ../cmd/posix-oneshot --storage_dir=${LOG} --entries="${LEAF}"
  size=$(sed -n '2 p' ${LOG}/checkpoint)
  cp ${LOG}/checkpoint ${LOG}/checkpoint.${size}
done

rm ${LEAF}
