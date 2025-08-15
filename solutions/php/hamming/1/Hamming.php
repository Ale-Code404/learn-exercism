<?php

function distance(string $strandA, string $strandB): int
{
    if(strlen($strandA) !== strlen($strandB)){
        throw new InvalidArgumentException('DNA strands must be of equal length.');
    }

    $differences = 0;
    foreach(str_split($strandA) as $index => $strandASingular) {
        if($strandASingular !== $strandB[$index]){
            $differences++;
        }
    }

    return $differences;
}
