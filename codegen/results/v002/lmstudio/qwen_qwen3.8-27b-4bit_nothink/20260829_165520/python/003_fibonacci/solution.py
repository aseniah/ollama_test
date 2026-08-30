import sys

def main():
    # Read the single integer argument from command line
    if len(sys.argv) < 2:
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        return
    
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to and including the largest one that doesn't exceed N
    # Sequence: 1, 1, 2, 3, 5, ...
    a, b = 1, 1
    if N >= 1:
        print(a)
        if N >= b and b > a:  # Avoid printing duplicate 1 if N is 1, but we need to print both 1s? 
            # Wait, the problem says "all Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N"
            # And "Start the sequence: 1, 1, 2, 3, 5, ..."
            # So if N=1, we should print 1, 1 (since both 1s are <= 1)
            # Let me re-read: "Print all Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N"
            # The Fibonacci sequence is 1, 1, 2, 3, 5, 8, ...
            # So if N=1, both 1s are <= 1, so we print both.
            # If N=2, we print 1, 1, 2.
            # If N=3, we print 1, 1, 2, 3.
            pass
    
    # Let me restructure
    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()