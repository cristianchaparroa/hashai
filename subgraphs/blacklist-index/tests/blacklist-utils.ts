import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import { Blacklisted } from "../generated/Blacklist/Blacklist"

export function createBlacklistedEvent(
  reportedAddress: Address,
  count: BigInt,
  category: BigInt,
  date: BigInt,
  comments: string,
  source: string
): Blacklisted {
  let blacklistedEvent = changetype<Blacklisted>(newMockEvent())

  blacklistedEvent.parameters = new Array()

  blacklistedEvent.parameters.push(
    new ethereum.EventParam(
      "reportedAddress",
      ethereum.Value.fromAddress(reportedAddress)
    )
  )
  blacklistedEvent.parameters.push(
    new ethereum.EventParam("count", ethereum.Value.fromUnsignedBigInt(count))
  )
  blacklistedEvent.parameters.push(
    new ethereum.EventParam(
      "category",
      ethereum.Value.fromUnsignedBigInt(category)
    )
  )
  blacklistedEvent.parameters.push(
    new ethereum.EventParam("date", ethereum.Value.fromUnsignedBigInt(date))
  )
  blacklistedEvent.parameters.push(
    new ethereum.EventParam("comments", ethereum.Value.fromString(comments))
  )
  blacklistedEvent.parameters.push(
    new ethereum.EventParam("source", ethereum.Value.fromString(source))
  )

  return blacklistedEvent
}
