import sys

def check_prime(num):
    if num <= 1:
        return False
    for i in range(2, int(num**0.5) + 1):
        if num % i == 0:
            return False
    return True

def print_primes_up_to_n(n):
    primes = []
    for number in range(2, n + 1):
        if check_prime(number):
            primes.append(number)
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        sys.exit(1)

    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer.")
        sys.exit(1)

    print_primes_up_to_n(N)