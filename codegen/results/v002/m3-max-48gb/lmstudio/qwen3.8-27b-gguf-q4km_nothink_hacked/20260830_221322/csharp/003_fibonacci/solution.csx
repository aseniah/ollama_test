using System;

if (Args.Count != 1)
{
    return;
}

if (!long.TryParse(Args[0], out long N))
{
    return;
}

if (N < 1)
{
    return;
}

// Generate Fibonacci numbers up to N
// Sequence: 1, 1, 2, 3, 5, ...
// We start with F1=1, F2=1, F3=2, F4=3, F5=5, ...

long a = 1;
long b = 1;

Console.WriteLine(a); // 1
if (N == 1) return;

Console.WriteLine(b); // 1

while (true)
{
    long next = a + b;
    if (next > N) break;
    Console.WriteLine(next);
    a = b;
    b = next;
}