using System;
using System.Linq;

if (Args.Length != 1 || !int.TryParse(Args[0], out int n))
{
    return;
}

if (n < 1)
{
    return;
}

long a = 1;
long b = 1;

// Print the first number if it's within the range
if (a <= n)
{
    Console.WriteLine(a);
}

// Generate and print subsequent numbers while they do not exceed n
while (b <= n)
{
    Console.WriteLine(b);
    long next = a + b;
    a = b;
    b = next;
}