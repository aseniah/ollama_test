using System;

if (Args.Count > 0) {
    int n = int.Parse(Args[0]);
    
    for (int i = 2; i <= n; i++) {
        bool isPrime = true;
        for (int d = 2; d * d <= i; d++) {
            if (i % d == 0) {
                isPrime = false;
                break;
            }
        }
        if (isPrime) {
            Console.WriteLine(i);
        }
    }
}