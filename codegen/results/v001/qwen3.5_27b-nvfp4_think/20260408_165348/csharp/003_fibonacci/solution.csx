using System;

if (Args.Count < 1) return;

long n = long.Parse(Args[0]);

if (n < 1) return;

long a = 1;
long b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
}