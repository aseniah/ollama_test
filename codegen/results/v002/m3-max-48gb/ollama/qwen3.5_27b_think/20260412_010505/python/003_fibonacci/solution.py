import sys

def main():
    # Check if a command line argument was provided
    if len(sys.argv) < 2:
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        # If argument is not a valid integer, exit silently
        return

    # If N < 1, print nothing
    if n < 1:
        return

    # Generate and print Fibonacci numbers up to and including the largest <= n
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()