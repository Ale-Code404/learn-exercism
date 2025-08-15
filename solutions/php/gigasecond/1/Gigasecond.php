<?php

function from(DateTimeImmutable $date): DateTimeImmutable
{
    $gSecond = (10**9);

    return $date->add(new DateInterval("PT{$gSecond}S"));
}
