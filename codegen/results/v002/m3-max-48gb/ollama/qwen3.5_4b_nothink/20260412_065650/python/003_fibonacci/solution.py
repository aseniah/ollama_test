import sys

if len(sys.argv) != 2:
    print("Usage: python fibonacci.py <N>", file=sys.stderr)
    sys.exit(1)

try:
    n = int(sys.argv[1])
except ValueError:
    print("Invalid input: N must be an integer", file=sys.stderr)
    sys.exit(1)

if n < 1:
    # Print nothing
    pass
else:
    a, b = 1, 1
    # Print first two 1s if n >= 1
    if n >= 1:
        print(a)
        print(b)
    
    # Generate subsequent Fibonacci numbers
    while True:
        c = a + b
        if c > n:
            break
        print(c)
        a, b = b, c