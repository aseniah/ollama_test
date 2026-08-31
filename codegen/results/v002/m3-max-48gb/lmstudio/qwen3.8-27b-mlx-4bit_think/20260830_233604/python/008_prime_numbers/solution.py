import sys

def main():
    if len(sys.argv) < 2:
        return
    N = int(sys.argv[1])

    if N < 2:
        return

    # Sieve of Eratosthenes
    sieve = [True] * (N + 1)
    sieve[0] = False
    sieve[1] = False
    for i in range(2, int(N ** 0.5) + 1):
        if sieve[i]:
            for j in range(i * i, N + 1, i):
                sieve[j] = False

    for i in range(2, N + 1):
        if sieve[i]:
            print(i)

main()