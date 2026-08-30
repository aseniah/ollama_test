import sys

def main():
    # Read N from command line arguments
    if len(sys.argv) != 2:
        print("Usage: python solution.py <N>")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: N must be an integer")
        sys.exit(1)
    
    # If N < 1, print nothing
    if n < 1:
        return
    
    # Start the Fibonacci sequence: 1, 1, 2, 3, 5, ...
    fib1 = 1
    fib2 = 1
    
    # We need to handle the case where N is exactly 0 or very small
    # Since the first two numbers are 1, if N < 1 we do nothing (already handled)
    # If N >= 1, we print the sequence elements <= N
    
    while True:
        # Check the current fib1
        if fib1 <= n:
            print(fib1)
        else:
            # fib1 > n, so stop
            break
        
        # Move to next Fibonacci number
        temp = fib2
        fib2 = fib1 + fib2
        fib1 = temp

if __name__ == "__main__":
    main()