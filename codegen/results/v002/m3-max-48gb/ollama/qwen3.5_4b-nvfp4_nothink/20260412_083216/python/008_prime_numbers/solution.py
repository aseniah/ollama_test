import sys

def is_prime(num):
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    
    # Check odd divisors from 3 up to sqrt(num)
    for i in range(3, int(num ** 0.5) + 1, 2):
        if num % i == 0:
            return False
    return True

def main():
    try:
        # Accept single integer argument from command line
        if len(sys.argv) != 2:
            print("Usage: python script.py <N>")
            sys.exit(1)
        
        N = int(sys.argv[1])
    except ValueError:
        print("Invalid input: Please provide a valid integer.")
        sys.exit(1)

    # Print all prime numbers up to and including N, one per line
    count = 0
    for num in range(2, N + 1):
        if is_prime(num):
            print(num)
            count += 1
    
    # If there are no primes (N < 2), print nothing (which happens automatically)

if __name__ == "__main__":
    main()