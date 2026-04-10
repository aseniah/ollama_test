import sys

def get_primes(n):
    """
    Returns a list of all prime numbers up to and including n
    using the Sieve of Eratosthenes.
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
                
    # Collect all indices that are still marked as True
    return [p for p in range(n + 1) if is_prime[p]]

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        # If no argument or too many, we exit silently or with error.
        # Based on requirements, we just need to handle the input N.
        return

    try:
        # Parse the command line argument to an integer
        n = int(sys.argv[1])
        
        # Get the list of primes
        primes = get_primes(n)
        
        # Print each prime on a new line
        for prime in primes:
            print(prime)
            
    except ValueError:
        # If the argument is not a valid integer, exit
        pass

if __name__ == "__main__":
    main()