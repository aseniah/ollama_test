import sys

if len(sys.argv) != 2:
    sys.exit(1)

try:
    n = int(sys.argv[1])
except ValueError:
    sys.exit(1)

if n < 1:
    sys.exit(0)

# Generate and print Fibonacci numbers up to and including N
a, b = 1, 1
while a <= n:
    print(a)
    a, b = b, a + b