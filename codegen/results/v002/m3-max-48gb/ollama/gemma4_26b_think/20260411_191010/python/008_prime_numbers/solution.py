import sys

def print_primes_up_to(n):
    """
    Uses the Sieve of Eratosthenes algorithm to find all prime numbers
    up to and including n, and prints them one per line.
    """
    if n < 2:
        # If N is less than 2, there are no prime numbers.
        # Per instructions, we print nothing.
        return

    # Initialize a boolean array "is_prime" of size n+1
    # We assume all numbers are prime initially.
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False  # 0 and 1 are not prime numbers

    # Start sieving from the first prime number, 2
    for p in range(2, int(n**0.5)_ + 1):
        if is_prime[p]:
            # If is_prime[p] is true, mark all multiples of p starting from p*p
            # as false. We start at p*p because smaller multiples have
            # already been marked by smaller prime factors.
            for i in range(p * p, n + 1, p):
                is_prime[i] = False

    # Iterate through the boolean array and print the indices that are True
    for number in range(2, n + 1):
        if is_prime[number]:
            print(number)

def main():
    # Check if an argument was provided via command line
    if len(sys.argv) != 2:
        # If the argument count is incorrect, we exit silently or handle as error.
        # The prompt implies a single integer argument will be provided.
        return

    try:
        # Attempt to parse the argument as an integer
        n = int(sys.argv[1])
        print_primes_up_to(n)
    except ValueError:
        # If the argument is not a valid integer, we exit without printing.
        pass

if __name__ == "__main__":
    main()