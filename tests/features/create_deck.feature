Feature: Create a new Deck

  Scenario: (01) Create 52 deck that is not shuffled (Happy Path)
    When a user creates a full deck that is not shuffled
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | false    | 52        |

  Scenario: (02) Create 52 deck that is shuffled (Happy Path)
    When a user creates a full deck that is shuffled
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | true     | 52        |

  Scenario: (03) Create partial deck that is not shuffled (Happy Path)
    When a user creates a partial deck that is not shuffled with the following cards:
      | AC |
      | 5S |
      | JD |
      | KH |
      | 2C |
    Then the user should receive a deck ID and the following results:
      | shuffled | remaining |
      | false    | 5         |
