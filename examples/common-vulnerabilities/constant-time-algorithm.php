<?php

function ConstantTimeEqualsBuiltin(string $expected, string $actual): bool
{
    return hash_equals($expected, $actual);
}
