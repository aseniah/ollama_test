import sys

def main():
    N = int(sys.argv[1])

    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()