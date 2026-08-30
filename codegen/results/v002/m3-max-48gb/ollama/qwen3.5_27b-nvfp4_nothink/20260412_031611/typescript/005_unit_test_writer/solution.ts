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
  { 
    description: "Quantity below 10 (no discount)", 
    price: 10, 
    quantity: 9, 
    expected: 90 // 10 * 9
  },
  { 
    description: "Boundary quantity 10 (9% discount)", 
    price: 10, 
    quantity: 10, 
    expected: 90 // 10 * 10 * 0.9
  },
  { 
    description: "Quantity just below 50 (9% discount)", 
    price: 20, 
    quantity: 49, 
    expected: 882 // 20 * 49 * 0.9 = 980 * 0.9
  },
  { 
    description: "Boundary quantity 50 (25% discount)", 
    price: 100, 
    quantity: 50, 
    expected: 3750 // 100 * 50 * 0.75
  },
  { 
    description: "Quantity above 50 (25% discount)", 
    price: 8, 
    quantity: 60, 
    expected: 360 // 8 * 60 * 0.75 = 480 * 0.75
  }
];

function runTests() {
  let passedCount = 0;
  
  for (const test of testCases) {
    const result = calculateDiscount(test.price, test.quantity);
    // Using a small epsilon for floating point comparison safety, though inputs here result in exact decimals
    const tolerance = 1e-9; 
    
    if (Math.abs(result - test.expected) < tolerance) {
      console.log(`PASS: ${test.description}`);
      passedCount++;
    } else {
      console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
  }

  console.log(`\nTotal: ${passedCount}/${testCases.length} tests passed.`);
}

runTests();