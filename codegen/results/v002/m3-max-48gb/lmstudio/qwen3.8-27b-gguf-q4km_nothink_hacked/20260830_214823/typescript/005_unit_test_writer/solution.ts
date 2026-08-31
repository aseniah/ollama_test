function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

const testCases: TestCase[] = [
    {
        description: "quantity 9 (below 10, no discount)",
        price: 10,
        quantity: 9,
        expected: 10 * 9
    },
    {
        description: "quantity 10 (boundary, 10% discount)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9
    },
    {
        description: "quantity 49 (below 50, 10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9
    },
    {
        description: "quantity 50 (boundary, 25% discount)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75
    },
];

let allPassed = true;

for (const tc of testCases) {
    const got = calculateDiscount(tc.price, tc.quantity);
    const expected = tc.expected;
    
    // Use a small epsilon for floating point comparison
    const epsilon = 1e-10;
    if (Math.abs(got - expected) < epsilon) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${expected}, got: ${got})`);
        allPassed = false;
    }
}

process.exit(allPassed ? 0 : 1);