import sys

def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    for i in range(3, int(n ** 0.5) + 1, 2):
        if n % i == 0:
            return False
    return True

def main():
    data = sys.argv
    if len(data) < 2:
        return
    try:
        n = int(data[1])
    except ValueError:
        return
    
    primes = []
    for num in range(2, n + 1):
        if is_prime(num):
            primes.append(num)
    
    for p in primes:
        print(p)

if __name__ == "__main__":
    main()