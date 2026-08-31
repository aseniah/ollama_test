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

function runTests() {
    const tests: TestCase[] = [
        {
            price: 10,
            quantity: 9,
            expected: 90, // 10 * 9 = 90 (no discount)
            description: "Quantity 9 (boundary: just below 10, no discount)"
        },
        {
            price: 10,
            quantity: 10,
            expected: 90, // 10 * 10 * 0.9 = 90 (10% discount)
            description: "Quantity 10 (boundary: starts 10% discount)"
        },
        {
            price: 10,
            quantity: 49,
            expected: 441, // 10 * 49 * 0.9 = 441 (10% discount)
            description: "Quantity 49 (boundary: just below 50, 10% discount)"
        },
        {
            price: 10,
            quantity: 50,
            expected: 375, // 10 * 50 * 0.75 = 375 (25% discount)
            description: "Quantity 50 (boundary: starts 25% discount)"
        }
    ];

    tests.forEach(test => {
        const result = calculateDiscount(test.price, test.quantity);
        
        // Use a small epsilon for floating point comparison
        if (Math.abs(result - test.expected) < 0.0001) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        }
    });
}

runTests();