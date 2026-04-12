import sys

def main():
    # Check if a command line argument is provided
    if len(sys.argv) != 2:
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        return

    if n < 1:
        return

    a, b = 1, 1

    # Print the first number if it doesn't exceed N
    while a <= n:
        print(a)
        # Move to the next Fibonacci number
        a, b = b, a + b

if __name__ == "__main__":
    main()