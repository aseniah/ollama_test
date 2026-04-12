import sys

def main():
    # Get the integer argument from command line
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Usage: python script.py <integer>")
        return

    # If N < 1, print nothing
    if n < 1:
        return

    # Generate Fibonacci numbers up to N
    a, b = 1, 1
    fib_count = [1]

    while b <= n:
        a, b = b, a + b
        if a > n:
            break
        fib_count.append(a)
        fib_count.append(b) if b <= n else None

    # Print each Fibonacci number on its own line
    for num in fib_count:
        print(num)

if __name__ == "__main__":
    main()