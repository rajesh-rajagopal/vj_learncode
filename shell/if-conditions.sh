#!/bin/bash

_source_version="a"
_current_version="b"
  if [[ -n ${_source_version} &&  -n ${_current_version} ]] 
  then 
    echo "true"
  else 
    echo "false"
  fi

