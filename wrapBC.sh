#!/bin/sh

iFile=$1 
oFile=$2

./makebarcodes.sh $1 > tmp1.html

ifile=$(realpath tmp1.html)

wkhtmltopdf $ifile $oFile

evince $ofile
