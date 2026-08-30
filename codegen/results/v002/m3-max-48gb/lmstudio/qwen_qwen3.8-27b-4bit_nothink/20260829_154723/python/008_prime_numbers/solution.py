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
    args = sys.argv[1:]
    if len(args) != 1:
        return
    try:
        N = int(args[0])
    except ValueError:
        return
    if N < 2:
        return
    for num in range(2, N + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()