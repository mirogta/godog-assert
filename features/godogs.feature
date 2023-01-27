Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

# good scenario with a correct assertion
Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

# good scenario with a wrong assertion - should fail on `assert`
Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 1
    Then there should be 0 remaining

# non-sensical scenario - should fail on `require`
Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 20
    Then there should be -8 remaining
