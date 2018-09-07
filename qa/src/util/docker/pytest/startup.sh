#! /bin/bash

if [ "$1" != "" ] && [ "$2" != "" ] ; then
    echo '' > __init__.py
    echo -e $1 > spec.py
    echo -e $2 >> spec.py
    py.test -v spec.py
else
    echo "fail! fail! fail!"
fi
