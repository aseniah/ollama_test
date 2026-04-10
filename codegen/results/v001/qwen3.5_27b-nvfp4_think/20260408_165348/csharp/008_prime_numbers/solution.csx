using System;

if (Args.Count > 0 && int.TryParse(Args[0], out var n))
{
    for (int i = 2; i <= n; i++)
    {
        bool isPrime = true;
        // Check for factors from 2 up to sqrt(i)
        // Use division to avoid potential overflow on j * j for large integers
        for (int j = 2; j <= i / j; j++)
        {
            if (i % j == 0)
            {
                isPrime = false;
                break;
            }
        }

        if (isPrime)
        {
            Console.WriteLine(i);
        }
    }
}