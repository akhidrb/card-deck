Feature: Open a Deck

  Scenario: (2.1) Open a deck (Happy Path)
    Given a user creates a partial deck that is not shuffled with the following cards:
      | AC |
      | 5S |
      | JD |
      | KH |
      | 2C |
    When the user requests to open the created deck
    Then the user should receive a deck with the following results:
      | shuffled | remaining | cards          |
      | false    | 5         | AC,5S,JD,KH,2C |
