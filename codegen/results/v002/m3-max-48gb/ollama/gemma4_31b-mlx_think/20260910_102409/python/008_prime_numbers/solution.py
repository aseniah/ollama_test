import sys

def print_primes(n):
    """
    Prints all prime numbers up to and including n using the Sieve of Eratosthenes.
    """
    if n < 2:
        return

    # Initialize a boolean array "is_prime[0..n]" 
    # and set all entries to True.
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False  # 0 and 1 are not prime numbers

    p = 2
    while (p * p <= n):
        # If is_prime[p] is not changed, then it is a prime
        if is_prime[p]:
            # Update all multiples of p starting from p*p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False
        p += 1

    # Print all prime numbers
    for p in range(2, n + 1):
        if is_prime[p]:
            print(p)

def main():
    # Check if the argument is provided
    if len(sys.argv) != 2:
        return

    try:
        # Convert the command line argument to an integer
        n = int(sys.argv[1])
        print_primes(n)
    except ValueError:
        # If the argument is not an integer, the program simply exits
        pass

if __name__ == "__main__":
    main()