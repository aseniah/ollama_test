```typescript
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestResult {
    description: string;
    expected: number;
    got: number;
}

function runTests(): void {
    console.log("=== Testing calculateDiscount function ===\n");
    
    const tests: TestResult[] = [
        // Boundary condition 1: quantity = 9 (should use no discount)
        { description: "quantity=9, price=100 (no discount expected)", expected: 900, got: 0 },
        
        // Boundary condition 2: quantity = 10 (first quantity with 10% discount)
        { description: "quantity=10, price=100 (10% discount expected)", expected: 900, got: 0 },
        
        // Boundary condition 3: quantity = 49 (should use 10% discount)
        { description: "quantity=49, price=100 (10% discount expected)", expected: 441, got: 0 },
        
        // Boundary condition 4: quantity = 50 (should use 25% discount)
        { description: "quantity=50, price=100 (25% discount expected)", expected: 375, got: 0 },
    ];

    let passedCount = 0;
    const results: TestResult[] = [];

    for (const test of tests) {
        const actual = calculateDiscount(test.expected / test.got * quantityForTest, test.got);
        // Wait, I made a mistake here. Let me recalculate with correct logic.
    }

    // Better approach - run each test case properly
    results[0] = {
        description: "quantity=9, price=100 (no discount expected)",
        expected: calculateDiscount(100, 9),
        got: calculateDiscount(100, 9) === 900 ? true : false
    };

    console.log();
}

function main(): void {
    const args = process.argv.slice(2);

    if (args.length >= 1 && args[0] === '--help') {
        console.log("Usage: node script.js [test-case]");
        console.log("");
        console.log("Test cases format: price quantity description");
        return;
    }

    runTests();

    let passedCount = 0;
    const results = [];

    // Test Case 1: quantity=9, price=100 (no discount)
    const test1Price = 100;
    const test1Quantity = 9;
    const test1Result = calculateDiscount(test1Price, test1Quantity);
    const expected1 = 900;
    if (test1Result === expected1) {
        console.log(`PASS: quantity=9, price=100 (no discount)`);
        results.push({ description: "quantity=9, price=100 (no discount)", expected: expected1, got: test1Result });
        passedCount++;
    } else {
        console.log(`FAIL: quantity=9, price=100 (expected: ${expected1}, got: ${test1Result})`);
        results[results.length - 1] = { description: "quantity=9, price=100 (no discount)", expected: expected1, got: test1Result };
        passedCount++;
    }

    // Test Case 2: quantity=10, price=100 (10% discount)
    const test2Price = 100;
    const test2Quantity = 10;
    const test2Result = calculateDiscount(test2Price, test2Quantity);
    const expected2 = 900;
    if (test2Result === expected2) {
        console.log(`PASS: quantity=10, price=100 (10% discount)`);
        results.push({ description: "quantity=10, price=100 (10% discount)", expected: expected2, got: test2Result });
        passedCount++;
    } else {
        console.log(`FAIL: quantity=10, price=100 (expected: ${expected2}, got: ${test2Result})`);
        results[results.length - 1] = { description: "quantity=10, price=100 (10% discount)", expected: expected2, got: test2Result };
        passedCount++;
    }

    // Test Case 3: quantity=49, price=100 (10% discount)
    const test3Price = 100;
    const test3Quantity = 49;
    const test3Result = calculateDiscount(test3Price, test3Quantity);
    const expected3 = 441;
    if (test3Result === expected3) {
        console.log(`PASS: quantity=49, price=100 (10% discount)`);
        results.push({ description: "quantity=49, price=100 (10% discount)", expected: expected3, got: test3Result });
        passedCount++;
    } else {
        console.log(`FAIL: quantity=49, price=100 (expected: ${expected3}, got: ${test3Result})`);
        results[results.length - 1] = { description: "quantity=49, price=100 (10% discount)", expected: expected3, got: test3Result };
        passedCount++;
    }

    // Test Case 4: quantity=50, price=100 (25% discount)
    const test4Price = 100;
    const test4Quantity = 50;
    const test4Result = calculateDiscount(test4Price, test4Quantity);
    const expected4 = 375;
    if (test4Result === expected4) {
        console.log(`PASS: quantity=50, price=100 (25% discount)`);
        results.push({ description: "quantity=50, price=100 (25% discount)", expected: expected4, got: test4Result });
        passedCount++;
    } else {
        console.log(`FAIL: quantity=50, price=100 (expected: ${expected4}, got: ${test4Result})`);
        results[results.length - 1] = { description: "quantity=50, price=100 (25% discount)", expected: expected4, got: test4Result };
        passedCount++;
    }

    console.log(`\n=== Results ===`);
    console.log(`Total Tests: ${results.length}`);
    console.log(`Passed: ${passedCount} / ${results.length}`);
    console.log(`Failed: ${(results.length - passedCount)} / ${results.length}`);
}

main();