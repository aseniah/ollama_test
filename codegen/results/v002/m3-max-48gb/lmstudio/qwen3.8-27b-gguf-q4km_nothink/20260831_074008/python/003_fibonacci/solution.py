import sys

def main():
    try:
        n = int(sys.argv[1])
    except (IndexError, ValueError):
        return

    if n < 1:
        return

    a, b = 1, 1
    # Print first Fibonacci number
    print(a)
    if b <= n:
        print(b)
    else:
        return

    a, b = b, a + b
    while b <= n:
        print(b)
        a, b = b, a + b

if __name__ == "__main__":
    main()