import sys

def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        return

    try:
        n = int(sys.argv[1])
        if n < 2:
            return
    except ValueError:
        print("Please provide a valid integer.")
        return

    def is_prime(num):
        if num < 2:
            return False
        if num == 2:
            return True
        if num % 2 == 0:
            return True
        for i in range(3, int(num**0.5) + 1, 2):
            if num % i == 0:
                return False
        return True

    primes = []
    for num in range(2, n + 1):
        if is_prime(num):
            primes.append(num)

    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()