function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTests(testName: string, price: number, expected: number, quantity: number): void {
    const result = calculateDiscount(price, quantity);
    if (result === expected) {
        console.log(`PASS: ${testName}`);
    } else {
        console.log(`FAIL: ${testName} (expected: ${expected}, got: ${result})`);
    }
}

// Test cases
runTests("Test Case 1: Quantity 9", 100, 90, 9);
runTests("Test Case 2: Quantity 10", 100, 100, 10);
runTests("Test Case 3: Quantity 49", 100, 445, 49);
runTests("Test Case 4: Quantity 50", 100, 500, 50);