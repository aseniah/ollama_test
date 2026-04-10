import sys
from sys import argv

def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    
    for i in range(3, int(n**0.5) + 1, 2):
        if n % i == 0:
            return False
    
    return True

def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py N", file=sys.stderr)
        sys.exit(1)

    try:
        n = int(sys.argv[1])
    except ValueError:
        print(f"Invalid input: {sys.argv[1]}", file=sys.stderr)
        sys.exit(1)

    if n < 2:
        return

    for i in range(2, n + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()