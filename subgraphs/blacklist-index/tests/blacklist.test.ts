import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { Address, BigInt } from "@graphprotocol/graph-ts"
import { Blacklisted } from "../generated/schema"
import { Blacklisted as BlacklistedEvent } from "../generated/Blacklist/Blacklist"
import { handleBlacklisted } from "../src/blacklist"
import { createBlacklistedEvent } from "./blacklist-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/developer/matchstick/#tests-structure-0-5-0

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let reportedAddress = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let count = BigInt.fromI32(234)
    let category = BigInt.fromI32(234)
    let date = BigInt.fromI32(234)
    let comments = "Example string value"
    let source = "Example string value"
    let newBlacklistedEvent = createBlacklistedEvent(
      reportedAddress,
      count,
      category,
      date,
      comments,
      source
    )
    handleBlacklisted(newBlacklistedEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/developer/matchstick/#write-a-unit-test

  test("Blacklisted created and stored", () => {
    assert.entityCount("Blacklisted", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "Blacklisted",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "reportedAddress",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "Blacklisted",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "count",
      "234"
    )
    assert.fieldEquals(
      "Blacklisted",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "category",
      "234"
    )
    assert.fieldEquals(
      "Blacklisted",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "date",
      "234"
    )
    assert.fieldEquals(
      "Blacklisted",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "comments",
      "Example string value"
    )
    assert.fieldEquals(
      "Blacklisted",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "source",
      "Example string value"
    )

    // More assert options:
    // https://thegraph.com/docs/en/developer/matchstick/#asserts
  })
})
