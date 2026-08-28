import sys

def print_fibonacci_up_to_n(n):
    if n < 1:
        return
    
    # Start the Fibonacci sequence
    a, b = 1, 1
    
    # Print the first number if it's within range
    if a <= n:
        print(a)
    
    # Print the second number if it's within range
    if b <= n:
        print(b)
    
    # Generate and print subsequent Fibonacci numbers
    while True:
        next_fib = a + b
        if next_fib > n:
            break
        print(next_fib)
        a, b = b, next_fib

# Read the command line argument
if len(sys.argv) != 2:
    print("Usage: python script.py <N>")
    sys.exit(1)

try:
    n = int(sys.argv[1])
    print_fibonacci_up_to_n(n)
except ValueError:
    print("Please provide a valid integer argument")
    sys.exit(1)