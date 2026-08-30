using System;

// Parse the command-line argument
int n = int.Parse(Args[0]);

// Check each number from 2 to n for primality
for (int i = 2; i <= n; i++)
{
    bool isPrime = true;
    
    // Check if i is divisible by any number from 2 to sqrt(i)
    for (int j = 2; j * j <= i; j++)
    {
        if (i % j == 0)
        {
            isPrime = false;
            break;
        }
    }
    
    // Print if prime
    if (isPrime)
    {
        Console.WriteLine(i);
    }
}