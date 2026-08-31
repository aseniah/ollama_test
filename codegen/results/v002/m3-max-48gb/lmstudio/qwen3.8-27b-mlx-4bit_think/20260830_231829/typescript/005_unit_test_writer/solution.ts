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
        description: "quantity 9 (below discount threshold)",
        price: 100,
        quantity: 9,
        expected: 900, // 100 * 9
    },
    {
        description: "quantity 10 (first discount tier starts)",
        price: 100,
        quantity: 10,
        expected: 900, // 100 * 10 * 0.9
    },
    {
        description: "quantity 49 (just below second discount tier)",
        price: 100,
        quantity: 49,
        expected: 4410, // 100 * 49 * 0.9
    },
    {
        description: "quantity 50 (second discount tier starts)",
        price: 100,
        quantity: 50,
        expected: 3750, // 100 * 50 * 0.75
    },
];

let allPassed = true;

for (const tc of testCases) {
    const got = calculateDiscount(tc.price, tc.quantity);
    // Use a small epsilon for floating point comparison
    const isPass = Math.abs(got - tc.expected) < 1e-9;
    if (isPass) {
        console.log(`PASS: ${tc.description}`);
    } else {
        allPassed = false;
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${got})`);
    }
}

if (allPassed) {
    console.log("\nAll tests passed.");
} else {
    console.log("\nSome tests failed.");
}