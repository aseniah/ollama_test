import sys

def get_primes_up_to(n):
    """
    Returns a list of prime numbers up to and including n 
    using the Sieve of Eratosthenes.
    """
    if n < 2:
        return []
    
    # Initialize a boolean list 'is_prime' where index represents the number
    # Assume all numbers are prime initially
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False  # 0 and 1 are not prime
    
    for p in range(2, int(n**0.5) + 1):
        if is_prime[p]:
            # Update all multiples of p starting from p*p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False
                
    # Collect all indices that are still marked as True
    return [p for p, prime in enumerate(is_prime) if prime]

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        # If no argument or too many arguments, we exit silently or with error
        # Based on instructions, we expect exactly one integer argument
        return

    try:
        # Parse the command line argument as an integer
        n = int(sys.argv[1])
        
        # Get the list of primes
        primes = get_primes_up_to(n)
        
        # Print each prime on a new line
        for prime in primes:
            print(prime)
            
    except ValueError:
        # If the argument is not a valid integer, exit
        pass

if __name__ == "__main__":
    main()