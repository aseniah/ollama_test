#r "System.Text.Json"
#r "Microsoft.CSharp"

if (Args.Count < 1)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

if (n < 2)
{
    // No primes less than 2, print nothing.
    return;
}

// Sieve of Eratosthenes to find all primes up to n
bool[] isComposite = new bool[n + 1];
int limit = (int)Math.Sqrt(n);

for (int i = 2; i <= limit; i++)
{
    if (!isComposite[i])
    {
        for (int j = i * i; j <= n; j += i)
        {
            isComposite[j] = true;
        }
    }
}

for (int i = 2; i <= n; i++)
{
    if (!isComposite[i])
    {
        Console.WriteLine(i);
    }
}