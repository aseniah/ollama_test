import sys

def main():
    # Check if exactly one command line argument is provided
    if len(sys.argv) != 2:
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        return

    # If N < 1, print nothing
    if n < 1:
        return

    # Generate and print Fibonacci numbers up to N
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()