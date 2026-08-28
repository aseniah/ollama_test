#r "netstandard"

using System;
using System.Collections.Generic;

// Check if argument is provided and valid
if (Args.Count == 0 || !int.TryParse(Args[0], out int n))
{
    return;
}

// If N < 1, print nothing
if (n < 1)
{
    return;
}

// Generate Fibonacci numbers up to N
int a = 1;
int b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    int next = a + b;
    a = b;
    b = next;
}