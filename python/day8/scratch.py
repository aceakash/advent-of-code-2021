needle = 'abcdu'
n = set(needle)

print(n)

haystack = 'gacdb'
h = set(haystack)

diff = n.difference(h)
print(diff, len(diff))