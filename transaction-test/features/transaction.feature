@transaction @stage
Feature: Transaction

  
  Scenario: test
    Given I save a single purchase with given values
    | amount | date         | description |
    | 500.00 | 20-01-2021   | blackfriday |
    Then I expect the status response to be 201
    And I expect the id of the purchase to be returned
    When I send a request to get the given purchase using the id returned and the the currency as Real
    Then I expect the status response to be 200
    And The returned response should have the given values
    | originalAmount | description | transactionDate | exchangeRate | currency | convertedAmount |
    | 500.00         | blackfriday | 20-01-2021      | 5.194        | Real     | 2597.00         |