import sys

def main():
    # Read the integer argument from command line
    if len(sys.argv) < 2:
        # No argument provided, print nothing
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        # Invalid argument, print nothing
        return
    
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to and including the largest that doesn't exceed N
    # Sequence starts: 1, 1, 2, 3, 5, ...
    a, b = 1, 1
    
    # First number
    if N >= 1:
        print(a)
    
    # Second number
    if N >= 1:
        print(b)
    
    # Generate subsequent Fibonacci numbers
    while True:
        c = a + b
        if c > N:
            break
        print(c)
        a, b = b, c

if __name__ == "__main__":
    main()