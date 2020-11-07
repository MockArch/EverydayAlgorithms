def is_isogram(string):
    _check = []
    for i in string.replace("-", "").replace(" ", "").lower():
        if i not in _check:
            _check.append(i)
        else:
            return False
    return True
