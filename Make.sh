#!/bin/bash

SCRIPT=$(readlink -f "$0")

SCRIPTPATH=$(dirname "$SCRIPT")
PARENTPATH=$SCRIPTPATH/

PS3='Please enter your choice: '

options=(
  "proto"
)

project=""
OS_LOCAL_FLAG="linux"
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  OS_LOCAL_FLAG="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  OS_LOCAL_FLAG="mac"
elif [[ "$OSTYPE" == "cygwin" ]]; then
  OS_LOCAL_FLAG="win"
elif [[ "$OSTYPE" == "msys" ]]; then
  OS_LOCAL_FLAG="win"
elif [[ "$OSTYPE" == "win32" ]]; then
  OS_LOCAL_FLAG="win"
elif [[ "$OSTYPE" == "freebsd"* ]]; then
  OS_LOCAL_FLAG="linux"
else
  echo "unknown os type"
  return 0
fi

echo "=========================================================================="
echo "work path: ["$PARENTPATH"]"
echo "you os is: ["$OS_LOCAL_FLAG"]"
echo "=========================================================================="


select opt in "${options[@]}";
do
  case $opt in
    "Quit")
      exit 0;;
    "proto")
      cd $PARENTPATH/shared/bin/$OS_LOCAL_FLAG
      echo "enter $PARENTPATH""shared/bin/$OS_LOCAL_FLAG"
      sh ./build.sh $PARENTPATH;;
  esac
  project=$opt
  break
done

echo $project "finish"