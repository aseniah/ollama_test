#r "System.Text.Json"

using System;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Collections.Generic;

if (Args.Count == 0) {
    // No arguments provided, print nothing
    return;
}

if (!int.TryParse(Args[0], out int n)) {
    // Invalid integer, print nothing
    return;
}

if (n < 2) {
    // No primes less than 2
    return;
}

// Sieve of Eratosthenes to find all primes up to N
bool[] isComposite = new bool[n + 1];
isComposite[0] = true;
isComposite[1] = true;

for (int i = 2; i * i <= n; i++) {
    if (!isComposite[i]) {
        for (int j = i * i; j <= n; j += i) {
            isComposite[j] = true;
        }
    }
}

// Print primes
for (int i = 2; i <= n; i++) {
    if (!isComposite[i]) {
        Console.WriteLine(i);
    }
}