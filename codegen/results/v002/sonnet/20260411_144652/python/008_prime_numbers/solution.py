import sys

def primes_up_to(n):
    if n < 2:
        return
    sieve = [True] * (n + 1)
    sieve[0] = sieve[1] = False
    for i in range(2, int(n**0.5) + 1):
        if sieve[i]:
            for j in range(i*i, n + 1, i):
                sieve[j] = False
    for i in range(2, n + 1):
        if sieve[i]:
            print(i)

if __name__ == "__main__":
    primes_up_to(int(sys.argv[1]))
