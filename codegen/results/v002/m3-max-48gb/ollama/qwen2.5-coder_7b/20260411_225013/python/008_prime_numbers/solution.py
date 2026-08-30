import sys

def is_prime(num):
    if num <= 1:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    for i in range(3, int(num**0.5) + 1, 2):
        if num % i == 0:
            return False
    return True

def print_primes_up_to_n(n):
    for number in range(2, n + 1):
        if is_prime(number):
            print(number)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        sys.exit(1)
    
    try:
        N = int(sys.argv[1])
        if N < 0:
            raise ValueError("N must be a non-negative integer")
        print_primes_up_to_n(N)
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)