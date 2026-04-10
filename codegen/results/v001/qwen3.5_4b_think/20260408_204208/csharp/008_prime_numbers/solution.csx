using System;
using System.Linq;

int n = int.Parse(Args[0]);

if (n < 2) {
    return;
}

// Use Sieve of Eratosthenes to find all primes
bool[] isPrime = new bool[n + 1];
for (int i = 2; i <= n; i++) {
    isPrime[i] = true;
}

for (int p = 2; p * p <= n; p++) {
    if (isPrime[p]) {
        for (int i = p * p; i <= n; i += p) {
            isPrime[i] = false;
        }
    }
}

for (int i = 2; i <= n; i++) {
    if (isPrime[i]) {
        Console.WriteLine(i);
    }
}