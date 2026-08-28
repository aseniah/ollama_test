using System;

if (Args.Length == 0)
{
    return;
}

if (!long.TryParse(Args[0], out long n) || n < 1)
{
    // Print nothing if N < 1 or is not a valid integer
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
    
    // Check for overflow in case N is very large
    if (current < 0) break; 
}