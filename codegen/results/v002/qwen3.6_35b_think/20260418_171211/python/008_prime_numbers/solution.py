import sys

def main():
    # Ensure exactly one argument is provided
    if len(sys.argv) != 2:
        return
        
    try:
        n = int(sys.argv[1])
    except ValueError:
        return
        
    # If N is less than 2, there are no prime numbers to print
    if n < 2:
        return
        
    # Sieve of Eratosthenes for efficient prime finding
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False
    
    # Mark non-prime numbers
    for i in range(2, int(n**0.5) + 1):
        if is_prime[i]:
            for j in range(i * i, n + 1, i):
                is_prime[j] = False
                
    # Print all primes up to and including N
    for i in range(2, n + 1):
        if is_prime[i]:
            print(i)

if __name__ == "__main__":
    main()