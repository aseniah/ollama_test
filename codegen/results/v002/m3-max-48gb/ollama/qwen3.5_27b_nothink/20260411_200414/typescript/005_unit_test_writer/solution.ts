// Function definition as provided
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner function
function runTest(description: string, price: number, quantity: number, expected: number): void {
    const result = calculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison safety, though integers/standard decimals usually work fine
    const epsilon = 0.000001;
    
    if (Math.abs(result - expected) < epsilon) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${result})`);
    }
}

// Main execution
function main(): void {
    // Test Case 1: Boundary condition (below 10) - No discount
    // Price: 10, Qty: 9 => 10 * 9 = 90
    runTest("Quantity 9 (No discount)", 10, 9, 90);

    // Test Case 2: Boundary condition (exactly 10) - 10% discount
    // Price: 10, Qty: 10 => 10 * 10 * 0.9 = 90
    runTest("Quantity 10 (10% discount)", 10, 10, 90);

    // Test Case 3: Boundary condition (below 50) - 10% discount
    // Price: 100, Qty: 49 => 100 * 49 * 0.9 = 4410
    runTest("Quantity 49 (10% discount)", 100, 49, 4410);

    // Test Case 4: Boundary condition (exactly 50) - 25% discount
    // Price: 100, Qty: 50 => 100 * 50 * 0.75 = 3750
    runTest("Quantity 50 (25% discount)", 100, 50, 3750);

    // Additional Test Case: Large quantity
    // Price: 10, Qty: 100 => 10 * 100 * 0.75 = 750
    runTest("Quantity 100 (25% discount)", 10, 100, 750);
}

main();