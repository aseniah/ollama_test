using System;

if (Args.Length == 0)
{
    return;
}

if (!long.TryParse(Args[0], out long n) || n < 1)
{
    // Print nothing if N < 1 or invalid input
    return;
}

long current = 1;
long next = 1;

while (current <= n)
{
    Console.WriteLine(current);
    
    long temp = current + next;
    current = next;
    next = temp;
}