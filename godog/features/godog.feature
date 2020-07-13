Feature: eat godogs
  In order to be happy,
  I need to eat godogs

  Scenario: Eat 4 out of 12
    Given there are 12 godogs
    When I eat 4
    Then there should be 7 remaining