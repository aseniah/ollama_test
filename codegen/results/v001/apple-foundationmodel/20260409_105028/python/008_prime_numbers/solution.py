import sys

def is_prime(num):
    if num <= 1:
        return False
    if num <= 3:
        return True
    if num % 2 == 0 or num % 3 == 0:
        return False
    i = 5
    while i * i <= num:
        if num % i == 0 or num % (i + 2) == 0:
            return False
        i += 6
    return True

def print_primes(n):
    for number in range(2, n + 1):
        if is_prime(number):
            print(number)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
    else:
        try:
            N = int(sys.argv[1])
            print_primes(N)
        except ValueError:
            print("Please provide a valid integer.")