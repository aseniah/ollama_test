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

function runTest(testCase: TestCase): void {
    const actual = calculateDiscount(testCase.price, testCase.quantity);
    const epsilon = 0.0001;
    
    if (Math.abs(actual - testCase.expected) < epsilon) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${actual})`);
    }
}

function main(): void {
    const testCases: TestCase[] = [
        {
            description: "Quantity 9 - No discount (< 10)",
            price: 100,
            quantity: 9,
            expected: 900  // 100 * 9 * 1.0
        },
        {
            description: "Quantity 10 - 10% discount (>= 10, < 50)",
            price: 100,
            quantity: 10,
            expected: 900  // 100 * 10 * 0.9
        },
        {
            description: "Quantity 49 - 10% discount (< 50)",
            price: 100,
            quantity: 49,
            expected: 4410  // 100 * 49 * 0.9
        },
        {
            description: "Quantity 50 - 25% discount (>= 50)",
            price: 100,
            quantity: 50,
            expected: 3750  // 100 * 50 * 0.75
        },
        {
            description: "Quantity 100 - 25% discount (>= 50)",
            price: 50,
            quantity: 100,
            expected: 3750  // 50 * 100 * 0.75
        },
        {
            description: "Quantity 1 - No discount (< 10)",
            price: 25,
            quantity: 1,
            expected: 25  // 25 * 1 * 1.0
        }
    ];

    console.log("Running calculateDiscount tests...\n");
    
    testCases.forEach(testCase => {
        runTest(testCase);
    });
    
    console.log("\nAll tests completed.");
}

main();