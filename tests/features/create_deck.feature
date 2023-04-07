Feature: Create a new Deck

  Scenario: (1.1) Create 52 deck that is not shuffled (Happy Path)
    When a user creates a full deck that is not shuffled
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | false    | 52        |

  Scenario: (1.2) Create 52 deck that is shuffled (Happy Path)
    When a user creates a full deck that is shuffled
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | true     | 52        |

  Scenario: (1.3) Create partial deck that is not shuffled (Happy Path)
    When a user creates a partial deck that is not shuffled with the following cards:
      | AC |
      | 5S |
      | JD |
      | KH |
      | 2C |
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | false    | 5         |

  Scenario: (1.4) Create partial deck that is not shuffled with invalid cards (Edge Case)
    When a user creates a partial deck that is not shuffled with the following cards:
      | AB |
      | 5T |
      | JD |
      | KH |
      | 2C |
    Then the user should receive a validation error
