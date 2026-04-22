using System;

if (Args.Count > 0)
{
    int n = int.Parse(Args[0]);
    if (n >= 2)
    {
        bool[] isComposite = new bool[n + 1];
        for (int i = 2; i * i <= n; i++)
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
    }
}