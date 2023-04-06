Feature: Open a Deck

  Scenario: (2.1) Open a deck (Happy Path)
    Given a user creates a partial deck that is not shuffled with the following cards:
      | AC |
      | 5S |
      | JD |
      | KH |
      | 2C |
    When the user requests to open the created deck
    Then the user should open a deck with following results:
      | shuffled | remaining |
      | false    | 5         |
    And the cards in the deck should be:
      | code | value | suit     |
      | AC   | ACE   | CLUBS    |
      | 5S   | 5     | SPADES   |
      | JD   | JACK  | DIAMONDS |
      | KH   | KING  | HEARTS   |
      | 2C   | 2     | CLUBS    |
