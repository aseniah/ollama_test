using System;
using System.Text.Json;

// Function to check if a number is prime
bool IsPrime(int num)
{
    if (num <= 1) return false;
    if (num <= 3) return true;
    if (num % 2 == 0 || num % 3 == 0) return false;
    
    for (int i = 5; i * i <= num; i += 6)
    {
        if (num % i == 0 || num % (i + 2) == 0) return false;
    }
    return true;
}

// Read a single integer argument from command line
int N = int.Parse(Args[0]);

// Print all prime numbers up to and including N
foreach (int i in Enumerable.Range(2, N + 1))
{
    if (IsPrime(i))
    {
        Console.WriteLine(i);
    }
}