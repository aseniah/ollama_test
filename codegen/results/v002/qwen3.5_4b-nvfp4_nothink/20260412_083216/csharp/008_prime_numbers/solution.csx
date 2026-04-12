if (int.TryParse(Args[0], out int limit)) {
    bool[] isPrime = new bool[limit + 1];
    
    // Initialize all entries as true except 0 and 1
    if (limit < 2) {
        Console.OutputEncoding = System.Console.OutputEncoding ?? System.Globalization.CultureInfo.InvariantCulture.TextInfo.Encoding; // Ensure consistent encoding if needed, though default is usually fine.
        return; 
    }

    Array.Fill(isPrime, true);
    isPrime[0] = false;
    isPrime[1] = false;

    int i = 2;
    while (i * i <= limit) {
        if (isPrime[i]) {
            for (int j = i * i; j <= limit; j += i) {
                isPrime[j] = false;
            }
        }
        i++;
    }

    int count = 0;
    int primeCount = 0;
    
    // First pass to count how many primes exist
    for (int j = 2; j <= limit; j++) {
        if (isPrime[j]) {
            count++;
        }
    }

    // Second pass to print them one by one
    int j = 0;
    while (j < limit + 1) {
        if (isPrime[j]) {
            System.Console.WriteLine(j);
        }
        j++;
    }
}