using System;

long n = long.Parse(Args[0]);

long a = 1;
long b = 1;

while (a <= n) {
    Console.WriteLine(a);
    long c = a + b;
    a = b;
    b = c;
}