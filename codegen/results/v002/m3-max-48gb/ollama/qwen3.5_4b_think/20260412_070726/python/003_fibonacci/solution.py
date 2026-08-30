import sys

def main():
    # Check if the correct number of arguments was provided
    if len(sys.argv) != 2:
        print("Usage: python program.py <N>")
        sys.exit(1)

    # Parse the integer argument
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: Please provide a valid integer argument")
        sys.exit(1)

    # If N < 1, print nothing
    if n < 1:
        return

    # Generate Fibonacci numbers and print those that do not exceed N
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()