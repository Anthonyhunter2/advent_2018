#!/bin/bash

totalList=( `cat ./input.txt` )
count=0
seenvalues=()
stop=True
iterations=0
while $stop; do
iterations=$(( $iterations + 1 ))
echo "$iterations"
for (( i=0; i < ${#totalList[@]};i++));do
    freq=${totalList[i]}
    checknum=$(echo "$count $freq" | bc)
    #  echo "CheckNum $checknum"
    if [[ $(grep "^$checknum$" checkfile.txt) ]]; then 
        echo "FOUND: $checknum"
        exit 0 
    else echo "$checknum" >> checkfile.txt
    fi

    # for x in ${seenvalues[@]}; do 
    #     checkagainst=$x
    #     if [[ $checknum -eq $checkagainst ]]; then
    #         echo "HI you can stop now: $checknum"
    #         stop=False
    #         exit 1
    #     fi
    # done
    # seenvalues+=($checknum)
     count=$(($checknum))

    done
done
echo $seenvalues
