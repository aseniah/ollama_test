import sys

N = int(sys.argv[1])

if N < 1:
    sys.exit()

a, b = 1, 1
print(a)

if N >= b:
    print(b)

while True:
    c = a + b
    if c > N:
        break
    print(c)
    a, b = b, c
