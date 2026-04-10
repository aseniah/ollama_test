import sys

def find_primes_up_to(n):
    """
    Uses the Sieve of Eratosthenes to find all prime numbers up to and including n.
    """
    if n < 2:
        return []

    # Create a boolean array "is_prime[0..n]" and initialize
    # all entries it as true. A value in is_prime[i] will
    # finally be false if i is Not a prime, else true.
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False

    for p in range(2, int(n**0.5) + 1):
        # If is_prime[p] is not changed, then it is a prime
        if is_prime[p]:
            # Update all multiples of p starting from p*p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False

    # Generate list of primes from the boolean array
    return [p for p, prime in enumerate(is_prime) if prime]

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        return

    try:
        # Parse the command line argument as an integer
        n = int(sys.argv[1])
        
        # Calculate primes
        primes = find_primes_up_to(n)
        
        # Print each prime on a new line
        for prime in primes:
            sys.stdout.write(str(prime) + '\n')
            
    except ValueError:
        # If the argument is not a valid integer, exit silently 
        # as per typical CLI utility behavior for invalid inputs
        pass

if __name__ == "__main__":
    main()