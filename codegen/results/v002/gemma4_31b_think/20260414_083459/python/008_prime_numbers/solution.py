import sys

def print_primes_up_to_n(n):
    """
    Prints all prime numbers up to and including n using the Sieve of Eratosthenes.
    """
    if n < 2:
        return

    # Initialize a boolean array "prime[0..n]" and set
    # all entries it to true. A value in prime[i] will
    # finally be false if i is Not a prime, else true.
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
    # Ensure a command line argument is provided
    if len(sys.argv) != 2:
        # If the user doesn't provide exactly one argument, 
        # the script exits silently or could print a usage message.
        sys.exit(1)

    try:
        # Convert the first command line argument to an integer
        n_val = int(sys.argv[1])
        print_primes_up_to_n(n_val)
    except ValueError:
        # If the argument is not a valid integer, exit silently
        sys.exit(1)