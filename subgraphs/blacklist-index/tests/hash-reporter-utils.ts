import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import { ReportCreated } from "../generated/HashReporter/HashReporter"

export function createReportCreatedEvent(
  reportedAddress: Address,
  count: BigInt,
  category: BigInt
): ReportCreated {
  let reportCreatedEvent = changetype<ReportCreated>(newMockEvent())

  reportCreatedEvent.parameters = new Array()

  reportCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "reportedAddress",
      ethereum.Value.fromAddress(reportedAddress)
    )
  )
  reportCreatedEvent.parameters.push(
    new ethereum.EventParam("count", ethereum.Value.fromUnsignedBigInt(count))
  )
  reportCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "category",
      ethereum.Value.fromUnsignedBigInt(category)
    )
  )

  return reportCreatedEvent
}
