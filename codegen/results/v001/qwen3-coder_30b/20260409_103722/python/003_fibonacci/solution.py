import sys

def main():
    # Read the command line argument
    if len(sys.argv) != 2:
        print("Usage: python fibonacci.py <N>")
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer argument")
        return
    
    # If N < 1, print nothing
    if N < 1:
        return
    
    # Generate and print Fibonacci numbers up to N
    a, b = 1, 1
    
    # Print first Fibonacci number if it's <= N
    if a <= N:
        print(a)
    
    # Print second Fibonacci number if it's <= N
    if b <= N:
        print(b)
    
    # Generate subsequent Fibonacci numbers
    while True:
        next_fib = a + b
        if next_fib > N:
            break
        print(next_fib)
        a, b = b, next_fib

if __name__ == "__main__":
    main()