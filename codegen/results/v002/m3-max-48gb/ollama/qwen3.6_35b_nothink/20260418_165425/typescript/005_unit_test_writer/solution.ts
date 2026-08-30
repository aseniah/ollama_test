function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(description: string, price: number, quantity: number, expected: number): void {
    const result = calculateDiscount(price, quantity);
    if (Math.abs(result - expected) < 1e-9) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${result})`);
    }
}

// Test 1: Below minimum threshold (quantity = 9) - no discount
runTest("No discount for quantity < 10 (quantity=9)", 10, 9, 90);

// Test 2: Boundary condition (quantity = 10) - 10% discount
runTest("10% discount for quantity >= 10 and < 50 (quantity=10)", 10, 10, 90);

// Test 3: Boundary condition (quantity = 49) - 10% discount
runTest("10% discount for quantity >= 10 and < 50 (quantity=49)", 10, 49, 441);

// Test 4: Boundary condition (quantity = 50) - 25% discount
runTest("25% discount for quantity >= 50 (quantity=50)", 10, 50, 375);

// Test 5: Above 50 threshold (quantity = 51) - 25% discount
runTest("25% discount for quantity >= 50 (quantity=51)", 10, 51, 382.5);

// Test 6: Zero quantity
runTest("Zero quantity returns 0", 10, 0, 0);