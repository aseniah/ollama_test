import sys

def print_primes_up_to(n):
    """
    Uses the Sieve of Eratosthenes to find and print all 
    prime numbers up to and including n.
    """
    if n < 2:
        return

    # Create a boolean array "is_prime[0..n]" and initialize
    # all entries it as true. A value in is_prime[i] will
    # finally be false if i is Not a prime, else true.
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False

    p = 2
    while (p * p <= n):
        # If is_prime[p] is not changed, then it is a prime
        if is_prime[p]:
            # Update all multiples of p starting from p*p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False
        p += 1

    # Print all indices that are still marked as True
    for p in range(2, n + 1):
        if is_prime[p]:
            print(p)

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        # If usage is incorrect, you might want to print usage info.
        # However, per instructions, we just need to handle the N input.
        return

    try:
        # Parse the argument from command line
        n = int(sys.argv[1])
        print_primes_up_to(n)
    except ValueError:
        # If the input is not a valid integer, exit silently or handle error
        pass

if __name__ == "__main__":
    main()