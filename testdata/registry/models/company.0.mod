name: Company
fields:
  UUID:
    type: UUID
    attributes:
      - immutable
      - mandatory
  ID:
    type: AutoIncrement
    attributes:
      - mandatory
  FoundedAt:
    type: Time
  Name:
    type: String
identifiers:
  primary: UUID
  record: ID
related:
  Address:
    type: HasOne
  ContactInfo:
    type: HasOne