import sys

def main():
    # Read N from command line argument
    if len(sys.argv) < 2:
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        return
    
    # If N < 1, print nothing
    if N < 1:
        return
    
    # Start with the first two Fibonacci numbers
    a, b = 1, 1
    
    # If N is 1, we have only one "1" to print (but the sequence is 1,1,2,3,...)
    # Actually, the sequence is defined as starting 1, 1, 2, 3, 5, ...
    # So if N >= 1, we print the first 1.
    # If N >= 1, print the first 1
    print(a)
    
    # If N >= 1, we've already handled the first 1.
    # Now, we need to handle the second 1 and onwards.
    # We can just iterate and print b if it's <= N, then continue.
    
    # Let's restructure: generate and print each fib number as we go.
    
    # Reset approach:
    a, b = 1, 1
    
    # Print first 1 if N >= 1 (already checked N >= 1)
    if N >= 1:
        print(a)
    
    # Now generate the rest
    # We need to continue as long as the next number is <= N
    # But we already printed the first 1. So start with the second 1.
    # Actually, let's do it cleanly:
    
    # Reset
    if N < 1:
        return
    
    a, b = 1, 1
    printed = []
    
    # First number
    if a <= N:
        print(a)
    
    # Second number
    if b <= N:
        print(b)
    
    # Continue with next numbers
    a, b = b, a + b  # Now a=1 (second), b=2 (third)
    while b <= N:
        print(b)
        a, b = b, a + b

if __name__ == "__main__":
    main()