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
        description: "quantity 9 (below 10, no discount)",
        price: 100,
        quantity: 9,
        expected: 100 * 9, // 900
    },
    {
        description: "quantity 10 (boundary, 10% discount)",
        price: 100,
        quantity: 10,
        expected: 100 * 10 * 0.9, // 900
    },
    {
        description: "quantity 49 (boundary, 10% discount)",
        price: 100,
        quantity: 49,
        expected: 100 * 49 * 0.9, // 4410
    },
    {
        description: "quantity 50 (boundary, 25% discount)",
        price: 100,
        quantity: 50,
        expected: 100 * 50 * 0.75, // 3750
    },
];

function runTests(tests: TestCase[]): void {
    let allPassed = true;

    for (const tc of tests) {
        const actual = calculateDiscount(tc.price, tc.quantity);

        // Use a small epsilon to handle floating-point comparison
        const tolerance = 1e-9;
        if (Math.abs(actual - tc.expected) < tolerance) {
            console.log(`PASS: ${tc.description}`);
        } else {
            allPassed = false;
            console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${actual})`);
        }
    }

    if (allPassed) {
        console.log("\nAll tests passed!");
    } else {
        console.log("\nSome tests failed.");
    }
}

runTests(testCases);