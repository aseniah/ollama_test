using System;

long N = long.Parse(Args[0]);

if (N < 1) return;

long a = 1, b = 1;
while (a <= N)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
    // Guard against overflow: if b already exceeds N, next iteration won't print
    if (b > N) break;
}