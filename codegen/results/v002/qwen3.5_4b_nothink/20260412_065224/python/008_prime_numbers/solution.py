import sys

def is_prime(num):
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    i = 3
    while i * i <= num:
        if num % i == 0:
            return False
        i += 2
    return True

def main():
    try:
        n = int(sys.argv[1])
    except IndexError:
        sys.exit(1)
    except ValueError:
        sys.exit(1)

    primes = []
    for num in range(2, n + 1):
        if is_prime(num):
            primes.append(num)

    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()