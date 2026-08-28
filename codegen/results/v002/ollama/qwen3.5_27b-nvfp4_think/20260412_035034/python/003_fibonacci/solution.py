import sys

if __name__ == "__main__":
    try:
        # Attempt to parse the integer argument from command line
        n = int(sys.argv[1])
    except (IndexError, ValueError):
        # If argument is missing or not an integer, do nothing
        pass
    else:
        # Initialize Fibonacci sequence starting with 1, 1
        a, b = 1, 1
        
        # Print numbers in the sequence that do not exceed n
        while a <= n:
            print(a)
            a, b = b, a + b