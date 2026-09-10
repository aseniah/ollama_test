if (Args.Count > 0 && int.TryParse(Args[0], out int N))
{
    if (N < 1)
    {
        // print nothing
    }
    else
    {
        int a = 1, b = 1;
        if (a <= N) Console.WriteLine(a);
        if (b <= N && a != b) Console.WriteLine(b);
        while (true)
        {
            int c = a + b;
            if (c > N) break;
            Console.WriteLine(c);
            a = b;
            b = c;
        }
    }
}