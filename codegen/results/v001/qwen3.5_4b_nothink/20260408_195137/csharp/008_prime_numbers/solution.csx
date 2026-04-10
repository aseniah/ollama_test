using System;

var N = int.Parse(Args[0]);

if (N < 2) {
    return;
}

var primes = new int[N + 1];
for (var i = 2; i <= N; i++) {
    primes[i] = 1;
}

for (var i = 2; i * i <= N; i++) {
    if (primes[i] == 1) {
        for (var j = i * i; j <= N; j += i) {
            primes[j] = 0;
        }
    }
}

for (var i = 2; i <= N; i++) {
    if (primes[i] == 1) {
        Console.WriteLine(i);
    }
}