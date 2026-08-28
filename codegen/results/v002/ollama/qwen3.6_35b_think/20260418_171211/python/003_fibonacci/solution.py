import sys

if len(sys.argv) < 2:
    sys.exit(0)

try:
    n = int(sys.argv[1])
except ValueError:
    sys.exit(0)

if n >= 1:
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b