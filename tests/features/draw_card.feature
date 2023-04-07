Feature: Draw Cards from Deck

  Scenario: (3.1) Draw Cards from a Deck (Happy Path)
    Given a user creates a partial deck that is not shuffled with the following cards:
      | AC |
      | 5S |
      | JD |
      | KH |
      | 2C |
      | 7S |
      | JC |
    When the user draws 3 card(s) from the deck
    Then the user should get the following cards:
      | code | value | suit     |
      | AC   | ACE   | CLUBS    |
      | 5S   | 5     | SPADES   |
      | JD   | JACK  | DIAMONDS |

  Scenario: (3.2) Draw Cards from a Deck with a count greater than number of cards in deck (Edge Case)
    Given a user creates a partial deck that is not shuffled with the following cards:
      | AC |
      | 5S |
      | JD |
      | KH |
      | 2C |
      | 7S |
      | JC |
    When the user draws 8 card(s) from the deck
    Then the user should receive a validation error
