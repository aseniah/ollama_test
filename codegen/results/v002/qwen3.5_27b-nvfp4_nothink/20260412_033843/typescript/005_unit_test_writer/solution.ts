function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

const testCases: TestCase[] = [
    {
        description: "Quantity less than 10 (no discount)",
        price: 10,
        quantity: 9,
        expected: 90 // 10 * 9
    },
    {
        description: "Boundary at quantity 10 (10% discount)",
        price: 10,
        quantity: 10,
        expected: 90 // 10 * 10 * 0.9
    },
    {
        description: "Quantity less than 50 but >= 10 (10% discount)",
        price: 20,
        quantity: 49,
        expected: 882 // 20 * 49 * 0.9
    },
    {
        description: "Boundary at quantity 50 (25% discount)",
        price: 20,
        quantity: 50,
        expected: 750 // 20 * 50 * 0.75
    }
];

function runTests() {
    let passedCount = 0;
    let failedCount = 0;

    for (const testCase of testCases) {
        const result = calculateDiscount(testCase.price, testCase.quantity);
        
        // Use a small epsilon for floating point comparison
        const epsilon = 0.000001;
        const isPassing = Math.abs(result - testCase.expected) < epsilon;

        if (isPassing) {
            console.log(`PASS: ${testCase.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
            failedCount++;
        }
    }

    // Print a summary
    console.log(`\nTotal: ${passedCount + failedCount} | Passed: ${passedCount} | Failed: ${failedCount}`);
    
    // Exit with non-zero code if any tests failed
    if (failedCount > 0) {
        process.exit(1);
    }
}

runTests();