import sys

try:
    N = int(sys.argv[1])
except IndexError:
    sys.exit(1)

if N < 1:
    sys.exit(0)

a, b = 1, 1
while b <= N:
    print(b)
    a, b = b, a + b