using System;
using System.Collections.Generic;

if (Args.Count < 1)
    return;

if (!int.TryParse(Args[0], out int n))
    return;

if (n < 1)
    return;

// Start sequence: 1, 1, 2, 3, 5, ...
// We need to print all Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N.

// Generate Fibonacci numbers and print them
long a = 1;
long b = 1;

if (a <= n)
    Console.WriteLine(a);

if (b <= n && b > a)
    Console.WriteLine(b);
else if (b <= n && b == a)
    // a and b are both 1, we already printed a, but b is also 1
    // Since the sequence starts 1, 1, we need to print the second 1 as well
    Console.WriteLine(b);

// Now generate subsequent Fibonacci numbers
long prev = 1; // This represents the first 1
long curr = 1; // This represents the second 1

while (true)
{
    long next = prev + curr;
    if (next > n)
        break;
    Console.WriteLine(next);
    prev = curr;
    curr = next;
}