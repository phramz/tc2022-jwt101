import hmac


# algorithm example: still insecure ...  do not use this
def constant_time_equals(a, b):
    if len(a) != len(b):
        return False

    result = 0
    for x, y in zip(a, b):
        result |= x ^ y
    return result == 0


# better user `hmac.compare_digest(a, b)`
def constant_time_equals_builtin(a, b):
    return hmac.compare_digest(a, b)
