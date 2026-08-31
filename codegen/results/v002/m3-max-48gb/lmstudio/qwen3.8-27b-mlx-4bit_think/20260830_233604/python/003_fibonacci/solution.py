import sys

N = int(sys.argv[1])

if N < 1:
    sys.exit()

a, b = 1, 1
while a <= N:
    print(a)
    a, b = b, a + b