using System;

if (Args.Count < 1)
    return;

if (!int.TryParse(Args[0], out int N))
    return;

if (N < 1)
    return;

// Generate Fibonacci numbers starting with 1, 1, 2, 3, 5, ...
// Print each number that does not exceed N
long a = 1, b = 1;
while (a <= N)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
}