import sys

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
        print("Usage: python prime.py <N>")
        sys.exit(1)
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Invalid input: N must be an integer")
        sys.exit(1)
    
    if N < 2:
        return
    
    for num in range(2, N + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()