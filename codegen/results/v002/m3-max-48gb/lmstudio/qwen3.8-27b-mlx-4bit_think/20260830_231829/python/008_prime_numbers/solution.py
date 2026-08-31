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
    N = int(sys.argv[1])
    for i in range(2, N + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()