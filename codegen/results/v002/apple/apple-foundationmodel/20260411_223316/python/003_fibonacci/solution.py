import sys

def print_fibonacci_up_to_n(n):
    if n < 1:
        print("")  # Print nothing if N is less than 1
        return

    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    # Check if the correct number of command line arguments is provided
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
        sys.exit(1)

    try:
        N = int(sys.argv[1])
        print_fibonacci_up_to_n(N)
    except ValueError:
        print("Please provide a valid integer for N.")