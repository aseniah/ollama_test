import sys

def main():
    # Parse N from command line arguments
    n = int(sys.argv[1])
    
    # If N is less than 2, there are no prime numbers to print
    if n < 2:
        return

    # Sieve of Eratosthenes to find primes up to N
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False
    
    for i in range(2, int(n**0.5) + 1):
        if is_prime[i]:
            for j in range(i*i, n + 1, i):
                is_prime[j] = False
                
    # Print all prime numbers up to N, one per line
    for i in range(2, n + 1):
        if is_prime[i]:
            print(i)

if __name__ == "__main__":
    main()