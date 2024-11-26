import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { Address, BigInt } from "@graphprotocol/graph-ts"
import { ReportCreated } from "../generated/schema"
import { ReportCreated as ReportCreatedEvent } from "../generated/HashReporter/HashReporter"
import { handleReportCreated } from "../src/hash-reporter"
import { createReportCreatedEvent } from "./hash-reporter-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/developer/matchstick/#tests-structure-0-5-0

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let reportedAddress = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let count = BigInt.fromI32(234)
    let category = BigInt.fromI32(234)
    let newReportCreatedEvent = createReportCreatedEvent(
      reportedAddress,
      count,
      category
    )
    handleReportCreated(newReportCreatedEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/developer/matchstick/#write-a-unit-test

  test("ReportCreated created and stored", () => {
    assert.entityCount("ReportCreated", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "ReportCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "reportedAddress",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "ReportCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "count",
      "234"
    )
    assert.fieldEquals(
      "ReportCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "category",
      "234"
    )

    // More assert options:
    // https://thegraph.com/docs/en/developer/matchstick/#asserts
  })
})
