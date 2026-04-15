import sys

def main():
    # Check command line arguments
    if len(sys.argv) != 2:
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