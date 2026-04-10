import sys

def main():
    if len(sys.argv) != 2:
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.exit(1)

    if n < 1:
        return

    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()