using System;

int n = int.Parse(Args[0]);

for (int i = 2; i <= n; i++) {
    bool isPrime = true;

    if (i > 2 && i % 2 == 0) {
        isPrime = false;
    } else {
        // Start checking odd divisors from 3 up to sqrt(i)
        int sqrt = Math.Sqrt(i);
        for (int d = 3; d <= sqrt; d += 2) {
            if (i % d == 0) {
                isPrime = false;
                break;
            }
        }
    }

    if (isPrime) {
        Console.WriteLine(i);
    }
}