function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTests() {
    // Test case 1: quantity = 9 (boundary - less than 10)
    const test1Price = 100;
    const test1Quantity = 9;
    const test1Expected = 900; // 100 * 9
    const test1Result = calculateDiscount(test1Price, test1Quantity);
    if (test1Result === test1Expected) {
        console.log(`PASS: quantity ${test1Quantity} (less than 10)`);
    } else {
        console.log(`FAIL: quantity ${test1Quantity} (less than 10) (expected: ${test1Expected}, got: ${test1Result})`);
    }

    // Test case 2: quantity = 10 (boundary - first discount tier)
    const test2Price = 100;
    const test2Quantity = 10;
    const test2Expected = 900; // 100 * 10 * 0.9
    const test2Result = calculateDiscount(test2Price, test2Quantity);
    if (test2Result === test2Expected) {
        console.log(`PASS: quantity ${test2Quantity} (first discount tier)`);
    } else {
        console.log(`FAIL: quantity ${test2Quantity} (first discount tier) (expected: ${test2Expected}, got: ${test2Result})`);
    }

    // Test case 3: quantity = 49 (boundary - less than 50)
    const test3Price = 100;
    const test3Quantity = 49;
    const test3Expected = 4410; // 100 * 49 * 0.9
    const test3Result = calculateDiscount(test3Price, test3Quantity);
    if (test3Result === test3Expected) {
        console.log(`PASS: quantity ${test3Quantity} (less than 50)`);
    } else {
        console.log(`FAIL: quantity ${test3Quantity} (less than 50) (expected: ${test3Expected}, got: ${test3Result})`);
    }

    // Test case 4: quantity = 50 (boundary - second discount tier)
    const test4Price = 100;
    const test4Quantity = 50;
    const test4Expected = 3750; // 100 * 50 * 0.75
    const test4Result = calculateDiscount(test4Price, test4Quantity);
    if (test4Result === test4Expected) {
        console.log(`PASS: quantity ${test4Quantity} (second discount tier)`);
    } else {
        console.log(`FAIL: quantity ${test4Quantity} (second discount tier) (expected: ${test4Expected}, got: ${test4Result})`);
    }

    // Test case 5: quantity = 100 (additional test)
    const test5Price = 50;
    const test5Quantity = 100;
    const test5Expected = 3750; // 50 * 100 * 0.75
    const test5Result = calculateDiscount(test5Price, test5Quantity);
    if (test5Result === test5Expected) {
        console.log(`PASS: quantity ${test5Quantity} (third discount tier)`);
    } else {
        console.log(`FAIL: quantity ${test5Quantity} (third discount tier) (expected: ${test5Expected}, got: ${test5Result})`);
    }
}

runTests();