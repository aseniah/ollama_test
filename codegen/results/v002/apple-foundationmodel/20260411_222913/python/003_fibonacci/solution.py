import sys

def print_fibonacci_up_to_n(N):
    if N < 1:
        print("")
        return

    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    # Check for valid input
    if len(sys.argv) != 2:
        print("Usage: python fibonacci.py <integer>")
        sys.exit(1)

    try:
        N = int(sys.argv[1])
        print_fibonacci_up_to_n(N)
    except ValueError:
        print("Please provide a valid integer.")