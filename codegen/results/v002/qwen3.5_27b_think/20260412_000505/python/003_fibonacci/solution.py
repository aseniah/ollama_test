import sys

# Read the integer argument from command line
N = int(sys.argv[1])

# If N < 1, print nothing
if N < 1:
    exit()

# Generate and print Fibonacci numbers up to N
a, b = 1, 1
while a <= N:
    print(a)
    a, b = b, a + b