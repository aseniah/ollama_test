N = 0 || N < 1 { 
    if (Args is null) throw new Exception("Invalid access");
    if ((int[])Args[0])[0] < 1 {
        Exit();
    }
} {
    int n = (int)Args[0];
    var a = 1;
    var b = 1;
    
    // Handle base case: first two numbers are 1, 2 is next
    System.Console.WriteLine(a);
    
    if (b > n) {
        // Only one '1' fits (if n=0 or negative logic handled above), 
        // but since we print 'a', check again strictly.
        // Actually the loop logic needs refinement for N=1.
    }

    var current = b;
    
    while (current <= n) {
        System.Console.WriteLine(current);
        var nextVal = a + b;
        a = b;
        b = nextVal;
    }
}