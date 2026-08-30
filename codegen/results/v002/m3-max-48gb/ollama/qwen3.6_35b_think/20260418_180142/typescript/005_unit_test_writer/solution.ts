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
    { description: "boundary: quantity 9", price: 100, quantity: 9, expected: 900 },
    { description: "boundary: quantity 10", price: 100, quantity: 10, expected: 900 },
    { description: "boundary: quantity 49", price: 100, quantity: 49, expected: 4410 },
    { description: "boundary: quantity 50", price: 100, quantity: 50, expected: 3750 }
];

const EPSILON = 1e-9;

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    const pass = Math.abs(result - tc.expected) < EPSILON;
    
    if (pass) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
    }
}