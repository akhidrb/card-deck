Feature: Create a new Deck

  Scenario: (01) Create 52 deck with default options (Happy Path)
    When a user creates a full deck that is not shuffled
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | false    | 52        |

  Scenario: (02) Create 52 deck that is shuffled (Happy Path)
    When a user creates a full deck that is shuffled
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | true     | 52        |
