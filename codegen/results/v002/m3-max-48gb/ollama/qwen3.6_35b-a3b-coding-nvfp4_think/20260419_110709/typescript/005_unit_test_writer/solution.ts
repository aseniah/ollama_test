function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(description: string, expected: number, actual: number): void {
    if (expected === actual) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
}

function main(): void {
    // Test case 1: Quantity 9 (Just below the first threshold)
    // Logic: quantity < 10
    // Expected: price * quantity
    runTest(
        "Quantity 9 (Single Tier)",
        100 * 9,
        calculateDiscount(100, 9)
    );

    // Test case 2: Quantity 10 (Start of the second tier)
    // Logic: quantity < 50
    // Expected: price * quantity * 0.9
    runTest(
        "Quantity 10 (Ten Percent Discount)",
        100 * 10 * 0.9,
        calculateDiscount(100, 10)
    );

    // Test case 3: Quantity 49 (Just below the second threshold)
    // Logic: quantity < 50
    // Expected: price * quantity * 0.9
    runTest(
        "Quantity 49 (Ten Percent Discount)",
        100 * 49 * 0.9,
        calculateDiscount(100, 49)
    );

    // Test case 4: Quantity 50 (Start of the third tier)
    // Logic: else
    // Expected: price * quantity * 0.75
    runTest(
        "Quantity 50 (Twenty Five Percent Discount)",
        100 * 50 * 0.75,
        calculateDiscount(100, 50)
    );
}

main();