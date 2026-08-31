function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases with boundary conditions
const testCases: Array<{ description: string; price: number; quantity: number; expected: number }> = [
    {
        description: "quantity 9 (no discount)",
        price: 10,
        quantity: 9,
        expected: 10 * 9, // 90
    },
    {
        description: "quantity 10 (10% discount)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9, // 90
    },
    {
        description: "quantity 49 (10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9, // 441
    },
    {
        description: "quantity 50 (25% discount)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75, // 375
    },
    {
        description: "quantity 0 (no discount)",
        price: 5,
        quantity: 0,
        expected: 5 * 0, // 0
    },
];

let allPassed = true;

for (const tc of testCases) {
    const got = calculateDiscount(tc.price, tc.quantity);
    const tolerance = 1e-10;
    if (Math.abs(got - tc.expected) < tolerance) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${got})`);
        allPassed = false;
    }
}

if (!allPassed) {
    process.exit(1);
}