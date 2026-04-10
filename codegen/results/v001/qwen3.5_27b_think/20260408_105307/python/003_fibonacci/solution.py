import sys

def main():
    # Check if argument is provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse integer argument
        n = int(sys.argv[1])
    except ValueError:
        # Handle non-integer input silently
        return

    # Constraint: If N < 1, print nothing
    if n < 1:
        return

    # Generate Fibonacci sequence
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()