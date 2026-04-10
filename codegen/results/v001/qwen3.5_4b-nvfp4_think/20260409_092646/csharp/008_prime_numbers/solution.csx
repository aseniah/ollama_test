using System;
using System.Collections.Generic;

int N = Args.Count > 0 ? int.Parse(Args[0]) : 2;

for (int i = 2; i <= N; i++) {
    bool isPrime = true;
    
    if (i % 2 == 0 && i > 2) {
        isPrime = false;
    } else if (i % 2 != 0) {
        for (int j = 3; j * j <= i; j += 2) {
            if (i % j == 0) {
                isPrime = false;
                break;
            }
        }
    }
    
    if (isPrime) {
        Console.WriteLine(i);
    }
}