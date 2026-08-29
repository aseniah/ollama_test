using System;

if (Args.Count == 1 && int.TryParse(Args[0], out int n))
{
    if (n >= 2)
    {
        bool[] isPrime = new bool[n + 1];

        for (int i = 2; i <= n; i++)
        {
            isPrime[i] = true;
        }

        for (int i = 2; (long)i * i <= n; i++)
        {
            if (isPrime[i])
            {
                for (long j = (long)i * i; j <= n; j += i)
                {
                    isPrime[(int)j] = false;
                }
            }
        }

        for (int i = 2; i <= n; i++)
        {
            if (isPrime[i])
            {
                Console.WriteLine(i);
            }
        }
    }
}