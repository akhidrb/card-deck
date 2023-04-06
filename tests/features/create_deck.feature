Feature: Create a new Deck

  Scenario: Create 52 deck with default options
    When a user creates a full deck that is not shuffled
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | false    | 52        |
