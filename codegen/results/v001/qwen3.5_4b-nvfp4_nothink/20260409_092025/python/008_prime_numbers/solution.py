import sys

def is_prime(n):
    if n < 2:
        return False
    if n == 2 or n == 3:
        return True
    if n % 2 == 0 or n % 3 == 0:
        return False
    
    i = 5
    while i * i <= n:
        if n % i == 0 or n % (i + 2) == 0:
            return False
        i += 6
    return True

def main():
    try:
        if len(sys.argv) != 2:
            print("Usage: python script.py <N>")
            sys.exit(1)
        
        n = int(sys.argv[1])
        if n < 2:
            return
    except ValueError:
        print("Argument must be an integer.")
        sys.exit(1)

    primes = []
    for num in range(2, n + 1):
        if is_prime(num):
            primes.append(num)
    
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()