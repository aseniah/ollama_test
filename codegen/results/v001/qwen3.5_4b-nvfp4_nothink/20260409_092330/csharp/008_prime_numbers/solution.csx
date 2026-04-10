using System;
using System.Collections.Generic;

var n = int.Parse(Args[0]);

if (n < 2) {
    return;
}

// Sieve of Eratosthenes
var limit = Math.Max(n, 10); // ensure at least 10 to check up to n
bool[] isPrime = new bool[limit + 1];
Array.Fill(isPrime, true);
isPrime[0] = false;
if (limit >= 2) isPrime[1] = false;

var primes = new List<int>();

for (var i = 2; i <= limit; i++) {
    if (isPrime[i]) {
        primes.Add(i);
        // Mark multiples
        if ((long)(i * i) <= limit) {
            for (var j = i * i; j <= limit; j += i) {
                isPrime[j] = false;
            }
        }
    }
}

// Output primes up to N
if (primes.Count > 0 && primes[primes.Count - 1] < n) {
    var lastPrimerIndex = primes.IndexOf(primes.Last());
    for (var i = 0; i <= lastPrimerIndex; i++) {
        Console.WriteLine(primes[i]);
    }
}

void Main() => exit(0);
Main();