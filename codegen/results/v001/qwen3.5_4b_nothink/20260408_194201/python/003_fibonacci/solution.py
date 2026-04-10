import sys

def main():
    # Accept a single integer argument N from the command line
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: Argument must be an integer")
        sys.exit(1)
    
    # If N < 1, print nothing
    if n < 1:
        return
    
    # Generate Fibonacci numbers up to N
    a, b = 1, 1
    
    # Print the first two 1s (if N >= 1)
    if n >= 1:
        print(a)
        if n >= b:
            print(b)
        else:
            # If n is exactly 1, we've already printed it twice in the check above, 
            # but logically the sequence is 1, 1, 2... 
            # If n=1, we print 1 then stop because the next is 1 which is <= 1, then next is 2.
            # Let's restructure to avoid double printing logic confusion.
            pass
    
    # A cleaner loop approach:
    fib_sequence = [1, 1]
    while fib_sequence[-1] <= n:
        next_fib = fib_sequence[-1] + fib_sequence[-2]
        if next_fib <= n:
            fib_sequence.append(next_fib)
        else:
            break
    
    for num in fib_sequence:
        print(num)

if __name__ == "__main__":
    main()