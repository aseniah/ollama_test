import math
import sys

def main():
    if len(sys.argv) != 2:
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        return

    if n < 2:
        return

    sieve = bytearray(b"\x01") * (n + 1)
    sieve[0] = sieve[1] = 0

    limit = math.isqrt(n)
    for i in range(2, limit + 1):
        if sieve[i]:
            start = i * i
            sieve[start:n + 1:i] = b"\x00" * (((n - start) // i) + 1)

    primes = [str(i) for i in range(2, n + 1) if sieve[i]]
    if primes:
        sys.stdout.write("\n".join(primes) + "\n")

if __name__ == "__main__":
    main()