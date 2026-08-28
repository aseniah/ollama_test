import sys

def main():
    # Check if there is a command line argument
    if len(sys.argv) < 2:
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, do nothing (or handle as needed)
        return

    # If N < 1, print nothing
    if n < 1:
        return

    # Start the sequence 1, 1, 2...
    a, b = 1, 1
    
    # Loop and print while current Fibonacci number is less than or equal to N
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()