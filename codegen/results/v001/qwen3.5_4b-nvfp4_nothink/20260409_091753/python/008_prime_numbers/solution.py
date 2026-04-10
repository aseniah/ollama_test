import sys

def is_prime(num):
    if num < 2:
        return False
    for i in range(2, int(num ** 0.5) + 1):
        if num % i == 0:
            return False
    return True

if __name__ == "__main__":
    try:
        n = int(sys.argv[1])
    except IndexError:
        sys.stderr.write("Usage: python solution.py <integer>\n")
        sys.exit(1)
    except ValueError:
        sys.stderr.write("Argument must be an integer.\n")
        sys.exit(1)

    primes = [num for num in range(2, n + 1) if is_prime(num)]
    for p in primes:
        print(p)