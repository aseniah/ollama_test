using System;
using System.Collections.Generic;

int n = int.Parse(Args[0]);

if (n < 2) {
    return;
}

bool[] isPrime = new bool[n + 1];
for (int i = 2; i <= n; i++) {
    isPrime[i] = true;
}

for (int i = 2; i * i <= n; i++) {
    if (isPrime[i]) {
        for (int j = i * i; j <= n; j += i) {
            isPrime[j] = false;
        }
    }
}

for (int i = 2; i <= n; i++) {
    if (isPrime[i]) {
        Console.WriteLine(i);
    }
}