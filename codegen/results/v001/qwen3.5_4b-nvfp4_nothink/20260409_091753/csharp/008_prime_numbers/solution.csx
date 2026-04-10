using System;
using System.Collections.Generic;

Args is List<string> = new();
if (Args.Count > 0) {
    var limit = int.Parse(Args[0]);
} else {
    limit = 0;
}

var primes = new List<int>();

for (int n = 2; n <= limit; n++) {
    bool isPrime = true;
    for (int d = 2; d * d <= n; d++) {
        if (n % d == 0) {
            isPrime = false;
            break;
        }
    }
    if (isPrime) {
        primes.Add(n);
    }
}

if (primes.Count > 0) {
    Console.WriteLine(primes[0]);
    for (int i = 1; i < primes.Count; i++) {
        Console.WriteLine(primes[i]);
    }
}