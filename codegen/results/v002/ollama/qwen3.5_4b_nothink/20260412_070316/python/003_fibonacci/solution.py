import sys

def main():
    # Accept a single integer argument N
    if len(sys.argv) != 2:
        sys.stderr.write("Usage: python fibonacci.py <N>\n")
        sys.exit(1)

    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.stderr.write("Argument must be an integer.\n")
        sys.exit(1)

    # If N < 1, print nothing
    if n < 1:
        return

    a, b = 1, 1
    while b <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()