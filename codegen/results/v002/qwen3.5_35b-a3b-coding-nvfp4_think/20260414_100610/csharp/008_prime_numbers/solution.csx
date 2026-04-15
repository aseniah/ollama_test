using System;

// Check if we have exactly one argument
if (Args.Count != 1)
{
    return;
}

// Parse the argument
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// Function to check if a number is prime
bool IsPrime(int num)
{
    if (num < 2) return false;
    if (num == 2) return true;
    if (num % 2 == 0) return false;
    for (int i = 3; i * i <= num; i += 2)
    {
        if (num % i == 0) return false;
    }
    return true;
}

// Print all primes up to n
for (int i = 2; i <= n; i++)
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}