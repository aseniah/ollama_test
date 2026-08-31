import sys

def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    i = 3
    while i * i <= n:
        if n % i == 0:
            return False
        i += 2
    return True

def main():
    if len(sys.argv) != 2:
        return
    try:
        N = int(sys.argv[1])
    except ValueError:
        return
    for num in range(2, N + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()