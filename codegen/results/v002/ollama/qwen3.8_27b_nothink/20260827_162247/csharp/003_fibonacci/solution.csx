using System;
using System.Collections.Generic;

int N;
if (Args != null && Args.Count > 0 && int.TryParse(Args[0], out int n))
{
    N = n;
}
else
{
    Console.Error.WriteLine("Usage: provide a single integer argument");
    return;
}

if (N < 1)
{
    return;
}

// Generate Fibonacci numbers up to and including the largest one that does not exceed N
int a = 1;
int b = 1;

if (a <= N)
{
    Console.WriteLine(a);
}

if (b <= N)
{
    Console.WriteLine(b);
}

int next = a + b;
while (next <= N)
{
    Console.WriteLine(next);
    a = b;
    b = next;
    next = a + b;
}