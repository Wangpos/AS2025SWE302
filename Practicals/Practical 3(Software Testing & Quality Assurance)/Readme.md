# Practical 3: Software Testing & Quality Assurance

## Specification-Based Testing of the Shipping Fee Calculator

This project demonstrates how to use **specification-based (black-box) testing** to ensure a shipping fee calculator in Go works exactly as the business rules require. We do this without looking at the code's internals, just like a real-world user would.

---

## What Was Achieved

- **Systematic, thorough testing** using three industry-standard techniques:
  1. **Equivalence Partitioning**: Grouping inputs into partitions and testing one value from each group.
  2. **Boundary Value Analysis (BVA)**: Focusing on the edges of valid/invalid input ranges.
  3. **Decision Table Testing**: Ensuring all combinations of business rules are covered.
- **All important cases and tricky edge cases are tested.**
- **Tests are robust to code changes**: If the implementation changes but the rules stay the same, the tests are still valid.

---

## How It Was Achieved

### 1. Equivalence Partitioning

- **Weight**: Partitioned into too small (invalid), valid, and too large (invalid).
- **Zone**: Partitioned into valid ("Domestic", "International", "Express") and invalid (anything else).
- **Insured**: Both true and false tested.
- **Test**: One value from each partition is tested to represent the whole group.

### 2. Boundary Value Analysis

- **Tested weights at the exact boundaries**: 0 (invalid), just above 0 (valid), 50 (valid), just above 50 (invalid).
- **Why**: Bugs often hide at the edges of valid/invalid ranges.

### 3. Decision Table Testing

- **Mapped all combinations** of valid/invalid weights and zones, with and without insurance, to their expected outcomes.
- **Why**: Ensures no business rule combination is missed.

---

## Example Test Output

![alt text](<assets/go test cmd.png>)

![alt text](<assets/go test -v cmd.png>)

```
=== RUN   TestCalculateShippingFeeV2_EquivalencePartitioning
--- PASS: TestCalculateShippingFeeV2_EquivalencePartitioning (0.00s)
=== RUN   TestCalculateShippingFeeV2_BoundaryValueAnalysis
    --- PASS: TestCalculateShippingFeeV2_BoundaryValueAnalysis/Weight_at_lower_invalid_boundary (0.00s)
    --- PASS: TestCalculateShippingFeeV2_BoundaryValueAnalysis/Weight_just_above_lower_boundary (0.00s)
    --- PASS: TestCalculateShippingFeeV2_BoundaryValueAnalysis/Weight_at_upper_valid_boundary (0.00s)
    --- PASS: TestCalculateShippingFeeV2_BoundaryValueAnalysis/Weight_just_above_upper_boundary (0.00s)
=== RUN   TestCalculateShippingFeeV2_DecisionTable
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_1:_Weight_too_low (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_1:_Weight_too_high (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_2:_Domestic (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_2:_Domestic_Insured (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_3:_International (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_3:_International_Insured (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_4:_Express (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_4:_Express_Insured (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_5:_Invalid_Zone (0.00s)
    --- PASS: TestCalculateShippingFeeV2_DecisionTable/Rule_5:_Invalid_Zone_Insured (0.00s)
PASS
ok      shipping        0.179s
```

This output shows all tests passed successfully, confirming the function behaves as expected across all specified scenarios.

---

## In Simple Terms

- We tested all important types of input, not every possible value.
- We checked the tricky edges where bugs love to hide.
- We made sure every business rule and combination is covered.
- **All tests passed, so the function is reliable and ready for use!**

---

## Why This Matters

- **Efficient**: You don’t waste time testing every possible value, just the important ones.
- **Effective**: You catch real bugs, especially at the edges and in complex rule combinations.
- **Maintainable**: If the business rules change, you can update the tests easily.

---

_Completed as a part of SWE302 - Software Testing & Quality Assurance_

_This project demonstrates the application of systematic testing techniques to ensure software quality._
