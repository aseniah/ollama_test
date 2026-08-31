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
    data = sys.argv[1]
    try:
        N = int(data)
    except ValueError:
        return
    
    if N < 2:
        return
    
    for i in range(2, N + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()