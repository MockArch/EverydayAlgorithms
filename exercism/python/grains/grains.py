def square(number):

    if number <= 0 or number > 64:
        raise ValueError("Zero Is not allowed")
    else:
        return 2**(number - 1)


def total():

    return sum([2**i for i in range(64)])
