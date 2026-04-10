import sys

def fibonacci_up_to(n):
    if n < 1:
        return

    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
        sys.exit(1)

    try:
        N = int(sys.argv[1])
        fibonacci_up_to(N)
    except ValueError:
        print("Please provide a valid integer for N.")
        sys.exit(1)