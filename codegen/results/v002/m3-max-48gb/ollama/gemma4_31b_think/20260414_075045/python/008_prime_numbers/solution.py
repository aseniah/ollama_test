import sys

def print_primes(n):
    """
    Prints all prime numbers up to and including n, one per line.
    Uses the Sieve of Eratosthenes for efficiency.
    """
    if n < 2:
        return

    # Create a boolean array "prime[0..n]" and initialize
    # all entries it as true. A value in prime[i] will
    # finally be false if i is Not a prime, else true.
    prime = [True for _ in range(n + 1)]
    p = 2
    while (p * p <= n):
        # If prime[p] is not changed, then it is a prime
        if prime[p] == True:
            # Update all multiples of p
            for i in range(p * p, n + 1, p):
                prime[i] = False
        p += 1

    # Print all prime numbers
    for p in range(2, n + 1):
        if prime[p]:
            print(p)

if __name__ == "__main__":
    # Check if the argument N is provided
    if len(sys.argv) != 2:
        # The prompt specifies accepting a single integer argument.
        # If not provided, we exit silently or could print a usage message.
        sys.exit(1)

    try:
        # Convert the command line argument to an integer
        N = int(sys.argv[1])
        print_primes(N)
    except ValueError:
        # If the argument is not a valid integer, we exit.
        sys.exit(1)