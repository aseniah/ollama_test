import sys

def print_primes(n):
    """
    Prints all prime numbers up to and including n.
    Uses the Sieve of Eratosthenes algorithm for efficiency.
    """
    if n < 2:
        return

    # Initialize a boolean array "prime[0..n]" 
    # prime[i] will be True if i is a prime
    prime = [True for _ in range(n + 1)]
    p = 2
    while (p * p <= n):
        # If prime[p] is not changed, then it is a prime
        if (prime[p] == True):
            # Updating all multiples of p
            for i in range(p * p, n + 1, p):
                prime[i] = False
        p += 1

    # Print all prime numbers
    for p in range(2, n + 1):
        if prime[p]:
            print(p)

if __name__ == "__main__":
    # Check if the command line argument is provided
    if len(sys.argv) != 2:
        # If the user didn't provide exactly one argument, we exit silently 
        # or you can add a usage message here.
        sys.exit(0)

    try:
        # Parse the command line argument as an integer
        n_val = int(sys.argv[1])
        print_primes(n_val)
    except ValueError:
        # If the argument is not a valid integer, print nothing as per typical CLI tool behavior
        sys.exit(0)