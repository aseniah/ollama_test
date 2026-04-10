import sys

def main():
    if len(sys.argv) != 2:
        # Ignore invalid usage per problem statement (or handle as needed)
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        return

    if n < 1:
        return

    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()