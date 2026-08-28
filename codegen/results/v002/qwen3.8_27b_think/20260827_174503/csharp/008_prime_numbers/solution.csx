using System;
using System.Text;

if (Args.Count < 1)
    return;

if (!int.TryParse(Args[0], out int n))
    return;

if (n < 2)
    return;

// Sieve of Eratosthenes
bool[] isPrime = new bool[n + 1];
Array.Fill(isPrime, true);
isPrime[0] = false;
isPrime[1] = false;

for (int i = 2; (long)i * i <= n; i++)
{
    if (isPrime[i])
    {
        for (int j = i * i; j <= n; j += i)
            isPrime[j] = false;
    }
}

var sb = new StringBuilder();
for (int i = 2; i <= n; i++)
{
    if (isPrime[i])
        sb.Append(i).Append('\n');
}

Console.Write(sb.ToString());