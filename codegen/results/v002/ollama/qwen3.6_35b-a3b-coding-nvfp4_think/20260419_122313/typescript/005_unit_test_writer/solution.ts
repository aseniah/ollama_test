function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    price: number;
    quantity: number;
    expected: number;
    description: string;
}

const testCases: TestCase[] = [
    { price: 100, quantity: 9, expected: 900, description: "Quantity 9 (below first threshold)" },
    { price: 100, quantity: 10, expected: 900, description: "Quantity 10 (first threshold)" },
    { price: 100, quantity: 49, expected: 4410, description: "Quantity 49 (below second threshold)" },
    { price: 100, quantity: 50, expected: 3750, description: "Quantity 50 (second threshold)" },
    { price: 50, quantity: 15, expected: 675, description: "Quantity 15 (mid-range)" }
];

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    if (result === tc.expected) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
    }
}